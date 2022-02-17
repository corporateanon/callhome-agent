package consumer

import (
	"fmt"
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
		fmt.Printf("Error parsing message payload: %s\n", err)
		return
	}
	fmt.Printf("Received message %s from %d\n", payload.Text, payload.ChatID)
	if err := c.options.OnTextMessage(payload.Text); err != nil {
		fmt.Printf("An error occurred while handling message: %s\n", err)
	}
}

func (c *Consumer) Connect() error {
	opts := mqtt.NewClientOptions().
		AddBroker(c.options.BrokerHost).
		SetAutoReconnect(true).
		SetConnectRetry(true).
		SetConnectRetryInterval((time.Second * 5))

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	if token := client.Subscribe(c.options.MessageTopic, 0, c.handleMessage); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func NewConsumer(options *Options) IConsumer {
	return &Consumer{options: options}
}
