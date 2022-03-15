package processor

import (
	"encoding/json"
	"time"

	"github.com/onosproject/analytics/pkg/logger"
	"github.com/onosproject/analytics/pkg/messages"
)

func processAlarm(alarmJSON string) ([]byte, error) {
	var alarm messages.Alarm
	err := json.Unmarshal([]byte(alarmJSON), &alarm)
	if err != nil {
		return []byte{}, err
	}
	enrichAlarm(&alarm)
	if logger.IfDebug() {
		logger.Debug("Alarm after enrich: %v", alarm)
	}
	message, err := json.Marshal(alarm)
	return message, err
}

func enrichAlarm(alarm *messages.Alarm) {
	alarm.Timestamp = time.Now()
}
