# callhome-agent

Simple app that shows a simple message box on Linux and Windows desktop each time it receives a message in MQTT topic. For my personal use.

## Installation

## Configuration

Environment variables:

- `MQTT_HOST` an MQTT URL in `[host:port]` format
- `MQTT_TOPIC` the topic used for the communication

## Protocol

Just send a `{"text":"Hello world"}` JSON string to the topic and that's it.
