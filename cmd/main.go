package main

import (
	"log"
	"os"
	"time"

	"github.com/corporateanon/callhome/internal/consumer"
	"github.com/corporateanon/callhome/internal/messagebox"
)

type CommandArgs struct {
	Message string
}

func handleTextMessage(message string) error {
	return messagebox.NewMessageBox().ShowMessage(message)
}

func main() {
	consumer, err := consumer.NewConsumer(
		consumer.WithBrokerHost(os.Getenv("MQTT_HOST")),
		consumer.WithMessageTopic(os.Getenv("MQTT_TOPIC")),
		consumer.WithOnTextMessage(handleTextMessage),
	)

	if err != nil {
		log.Fatalf("Initialization error: %s\n", err)
	}

	if err := consumer.Connect(); err != nil {
		log.Fatalf("Failed to start due to error: %s\n", err)
	}

	for {
		time.Sleep(time.Second * 1e9)
	}
}
