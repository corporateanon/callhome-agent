package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/corporateanon/callhome/internal/messagebox"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MessagePayload struct {
	Text   string `json:"text"`
	ChatID int64  `json:"chatId"`
}

func parseMessagePayload(raw []byte) (*MessagePayload, error) {
	pl := MessagePayload{}
	err := json.Unmarshal(raw, &pl)
	if err != nil {
		return nil, err
	}
	return &pl, nil
}

type CommandArgs struct {
	Message string
}

func handleMessage(c mqtt.Client, m mqtt.Message) {
	payload, err := parseMessagePayload(m.Payload())
	if err != nil {
		fmt.Printf("Error parsing message payload: %s\n", err)
		return
	}
	fmt.Printf("Received message %s from %d\n", payload.Text, payload.ChatID)
	if err := messagebox.NewMessageBox().ShowMessage(payload.Text); err != nil {
		fmt.Printf("ShowMessage error %s\n", err)
		return
	}
}

func main() {
	host := os.Getenv("MQTT_HOST")
	topic := os.Getenv("MQTT_TOPIC")

	opts := mqtt.NewClientOptions().
		AddBroker(host).
		SetAutoReconnect(true).
		SetConnectRetry(true).
		SetConnectRetryInterval((time.Second * 5))

	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe(topic, 0, handleMessage); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	time.Sleep(1e9 * time.Second)

}
