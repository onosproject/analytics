package messages

import (
	"time"
)

type Alarm struct {
	Type      string    `json:"type"`
	Device    string    `json:"device,omitempty"`
	Message   string    `json:"message,omitempty"`
	Serverity uint      `json:"serverity,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (alarm Alarm) MessageType() MessageType {
	return ALARM
}
