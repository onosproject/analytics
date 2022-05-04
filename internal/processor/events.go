package processor

import (
	"encoding/json"
	"time"

	"github.com/onosproject/analytics/pkg/messages"
)

func processEvent(eventJSON string) ([]byte, error) {
	var event messages.Event
	err := json.Unmarshal([]byte(eventJSON), &event)
	if err != nil {
		return []byte{}, err
	}
	enrichEvent(&event)
	log.Debugf("Event after enrich: %v", event)
	message, err := json.Marshal(event)
	return message, err
}

func enrichEvent(event *messages.Event) {
	event.Timestamp = time.Now()
}
