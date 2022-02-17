package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"text/template"
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

	commandTemplate, err := template.New("command").Parse(os.Getenv("COMMAND"))
	if err != nil {
		fmt.Printf("Error parsing command template: %s\n", err)
		return
	}

	buf := bytes.NewBuffer([]byte{})
	if err := commandTemplate.Execute(buf, CommandArgs{Message: strconv.Quote(payload.Text)}); err != nil {
		fmt.Printf("Error executing command template: %s\n", err)
		return
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("c:\\windows\\system32\\cmd.exe", "/c", buf.String())
	} else {
		cmd = exec.Command("/bin/sh", "-c", buf.String())
	}

	stdout, err := cmd.Output()
	fmt.Printf("Running command %s\n", cmd.String())
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
