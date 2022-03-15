package messages

import (
	"time"
)

type Event struct {
	Type      string    `json:"type"`
	Device    string    `json:"device"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (event Event) MessageType() MessageType {
	return EVENT
}
