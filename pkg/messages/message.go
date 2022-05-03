package messages

import (
	"encoding/json"

	"github.com/onosproject/onos-lib-go/pkg/logging"
)

type MessageType int

const (
	ALARM MessageType = iota
	EVENT
	METRIC
)

var log = logging.GetLogger("messages")

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
	log.Debugf("GetModel(%s, %s)",
		getMessageTypeName(messageType), string(msgJson))
	switch messageType {
	case ALARM:
		var alarm Alarm
		err := json.Unmarshal(msgJson, &alarm)
		if err != nil {
			log.Errorf("Failed to Unmarshal %s, err:%v",
				string(msgJson), err)
			return nil, err
		}
		return alarm, nil
	case EVENT:
		var event Event
		err := json.Unmarshal(msgJson, &event)
		if err != nil {
			log.Errorf("Failed to Unmarshal %s, err:%v",
				string(msgJson), err)
			return nil, err
		}
		return event, nil
	case METRIC:
		var metric Metric
		err := json.Unmarshal(msgJson, &metric)
		if err != nil {
			log.Errorf("Failed to Unmarshal %s, err:%v",
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
	log.Debugf("GetJson(%v)\n", message)
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Errorf("json.Marshal for %v failed error: %v\n", message, err)
		return []byte{}, err
	}
	log.Debugf("%v.GetJson() returning %s\n", message, string(bytes))
	return bytes, nil
}
