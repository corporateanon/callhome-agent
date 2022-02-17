package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

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

func handleMessage(c mqtt.Client, m mqtt.Message) {
	payload, err := parseMessagePayload(m.Payload())
	if err != nil {
		fmt.Printf("Error parsing message payload: %s\n", err)
		return
	}
	fmt.Printf("Received message %s from %d\n", payload.Text, payload.ChatID)
	command := os.Getenv("COMMAND")
	fmt.Printf("Going to execute command %s\n", command)
	cmd := exec.Command(command, payload.Text)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing command %s\n", err)
		return
	}
	fmt.Printf("Command output: %s\n", stdout)
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
