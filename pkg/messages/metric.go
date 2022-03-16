package messages

import (
	"time"
)

/*
The Metric structure is meant to be a superset
of all metric message formats.  The omitempty annotation
will suppress marshalling unneeded fields
*/
type Metric struct {
	Device    string    `json:"device"`
	BytesIn   uint64    `json:"bytes_in,omitempty"`
	BytesOut  uint64    `json:"bytes_out,omitempty"`
	ErrorsIn  uint64    `json:"errors_in,omitempty"`
	ErrorsOut uint64    `json:"errors_out,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

/*
MessageType is required to satisfy the Message interface definition
*/
func (metric Metric) MessageType() MessageType {
	return METRIC
}
