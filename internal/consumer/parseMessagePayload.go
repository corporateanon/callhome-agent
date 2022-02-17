package consumer

import "encoding/json"

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
