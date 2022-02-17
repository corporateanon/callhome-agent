package main

import (
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
	opts := consumer.NewOptions().
		SetBrokerHost(os.Getenv("MQTT_HOST")).
		SetMessageTopic(os.Getenv("MQTT_TOPIC")).
		SetOnTextMessage(handleTextMessage)

	if err := consumer.NewConsumer(opts).Connect(); err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 1e9)
}
