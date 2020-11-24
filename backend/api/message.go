package api

import "encoding/json"

//go:generate stringer -type=Message

type Message struct {
	Action string
	Token  string
	Args   map[string]string
}

func ParseMessage(raw []byte) (Message, error) {
	var msg Message
	err := json.Unmarshal(raw, &msg)
	return msg, err
}
