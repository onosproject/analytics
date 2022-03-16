package messages

import (
	"encoding/json"

	"github.com/onosproject/analytics/pkg/logger"
)

type MessageType int

const (
	ALARM MessageType = iota
	EVENT
	METRIC
)

/*
Message interface enables single functions to be used to
process all types
*/
type Message interface {
	MessageType() MessageType
}

func getMessageTypeName(messageType MessageType) string {
	switch messageType {
	case ALARM:
		return "Alarm"
	case EVENT:
		return "Event"
	case METRIC:
		return "Metric"
	default:
		return "Unknown"
	}
}

/*
GetModel
takes Json and unmarshals into appropriate message model
*/
func GetModel(messageType MessageType, msgJson []byte) (Message, error) {
	if logger.IfDebug() {
		logger.Debug("GetModel(%s, %s)",
			getMessageTypeName(messageType), string(msgJson))
	}
	switch messageType {
	case ALARM:
		var alarm Alarm
		err := json.Unmarshal(msgJson, &alarm)
		if err != nil {
			logger.Error("Failed to Unmarshal %s, err:%v",
				string(msgJson), err)
			return nil, err
		}
		return alarm, nil
	case EVENT:
		var event Event
		err := json.Unmarshal(msgJson, &event)
		if err != nil {
			logger.Error("Failed to Unmarshal %s, err:%v",
				string(msgJson), err)
			return nil, err
		}
		return event, nil
	case METRIC:
		var metric Metric
		err := json.Unmarshal(msgJson, &metric)
		if err != nil {
			logger.Error("Failed to Unmarshal %s, err:%v",
				string(msgJson), err)
			return nil, err
		}
		return metric, nil
	}
	return nil, nil

}

/*
GetJson
converts message struct (event/alarm/metric) into json
byte array suitable for sending to kafka bus
*/
func GetJson(message Message) ([]byte, error) {
	if logger.IfDebug() {
		logger.Debug("GetJson(%v)\n", message)
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		logger.Error("json.Marshal for %v failed error: %v\n", message, err)
		return []byte{}, err
	}
	if logger.IfDebug() {
		logger.Debug("%v.GetJson() returning %s\n", message, string(bytes))
	}
	return bytes, nil
}
