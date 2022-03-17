package messages

import (
	"time"
)

/*
The Event structure is meant to be a superset
of all event message formats.  The omitempty annotation
will suppress marshalling unneeded fields
*/
type Event struct {
	Type      string    `json:"type"`
	Device    string    `json:"device,omitempty"`
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

/*
MessageType is required to satisfy the Message interface definition
*/
func (event Event) MessageType() MessageType {
	return EVENT
}
