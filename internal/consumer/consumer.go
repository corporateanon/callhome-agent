//go:generate go-options -prefix With Consumer

package consumer

import (
	"crypto/tls"
	"log"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type IConsumer interface {
	Connect() error
}

type Consumer struct {
	brokerHost    string
	messageTopic  string
	onTextMessage func(msg string) error
}

func (c *Consumer) handleMessage(client mqtt.Client, m mqtt.Message) {
	payload, err := parseMessagePayload(m.Payload())
	if err != nil {
		log.Printf("Error parsing message payload: %s\n", err)
		return
	}
	log.Printf("Received message %s from %d\n", payload.Text, payload.ChatID)
	if err := c.onTextMessage(payload.Text); err != nil {
		log.Printf("An error occurred while handling message: %s\n", err)
	}
}

func (c *Consumer) Connect() error {
	opts := mqtt.NewClientOptions().
		AddBroker(c.brokerHost).
		SetKeepAlive(10 * time.Second).
		SetConnectRetry(true).
		SetAutoReconnect(true).
		SetConnectRetryInterval(time.Second * 5).
		SetMaxReconnectInterval(time.Minute * 1).
		SetConnectionLostHandler(func(c mqtt.Client, err error) {
			log.Printf("Connection lost due to error %s\n", err)
		}).
		SetConnectionAttemptHandler(func(broker *url.URL, tlsCfg *tls.Config) *tls.Config {
			log.Println("Connecting...")
			return tlsCfg
		}).
		SetOnConnectHandler(func(client mqtt.Client) {
			log.Println("Connected")
			if token := client.Subscribe(c.messageTopic, 0, c.handleMessage); token.Wait() && token.Error() != nil {
				log.Printf("Could not subscribe to the topic due to error %s\n", token.Error())
			}
		})

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func NewConsumer(options ...Option) (IConsumer, error) {
	c, err := newConsumer(options...)
	return &c, err
}
