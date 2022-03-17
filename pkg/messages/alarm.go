package messages

import (
	"time"
)

/*
The Alarm structure is meant to be a superset
of all alarm message formats.  The omitempty annotation
will suppress marshalling unneeded fields
*/
type Alarm struct {
	Type      string    `json:"type"`
	Device    string    `json:"device,omitempty"`
	Message   string    `json:"message,omitempty"`
	Serverity uint      `json:"serverity,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

/*
MessageType is required to satisfy the Message interface definition
*/
func (alarm Alarm) MessageType() MessageType {
	return ALARM
}
