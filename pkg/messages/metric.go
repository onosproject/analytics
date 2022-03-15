package messages

import (
	"time"
)

type Metric struct {
	Device    string    `json:"device"`
	BytesIn   uint64    `json:"bytes_in"`
	BytesOut  uint64    `json:"bytes_out"`
	ErrorsIn  uint64    `json:"errors_in"`
	ErrorsOut uint64    `json:"errors_out"`
	Timestamp time.Time `json:"timestamp"`
}

func (metric Metric) MessageType() MessageType {
	return METRIC
}
