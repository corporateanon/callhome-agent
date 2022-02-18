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
	options *Options
}

func (c *Consumer) handleMessage(client mqtt.Client, m mqtt.Message) {
	payload, err := parseMessagePayload(m.Payload())
	if err != nil {
		log.Printf("Error parsing message payload: %s\n", err)
		return
	}
	log.Printf("Received message %s from %d\n", payload.Text, payload.ChatID)
	if err := c.options.onTextMessage(payload.Text); err != nil {
		log.Printf("An error occurred while handling message: %s\n", err)
	}
}

func (c *Consumer) Connect() error {
	opts := mqtt.NewClientOptions().
		AddBroker(c.options.brokerHost).
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
		SetOnConnectHandler(func(c mqtt.Client) {
			log.Println("Connected")
		})

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	if token := client.Subscribe(c.options.messageTopic, 0, c.handleMessage); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func NewConsumer(options *Options) IConsumer {
	return &Consumer{options: options}
}
