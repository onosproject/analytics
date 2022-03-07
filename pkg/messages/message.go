package messages

import (
	"encoding/json"
	"log"
)

type MessageType int

const (
	ALARM MessageType = iota
	EVENT
	METRIC
)

type Message interface {
	MessageType() MessageType
}

func GetModel(messageType MessageType, msgJson []byte) (Message, error) {
	switch messageType {
	case ALARM:
		var alarm Alarm
		err := json.Unmarshal(msgJson, &alarm)
		if err != nil {
			log.Printf("Failed to Unmashal %s err: %v\n",
				string(msgJson), err)
			return nil, err
		}
		return alarm, nil
	case EVENT:
		var event Event
		err := json.Unmarshal(msgJson, &event)
		if err != nil {
			log.Printf("Failed to Unmashal %s err: %v\n",
				string(msgJson), err)
			return nil, err
		}
		return event, nil
	case METRIC:
		var metric Metric
		err := json.Unmarshal(msgJson, &metric)
		if err != nil {
			log.Printf("Failed to Unmashal %s err: %v\n",
				string(msgJson), err)
			return nil, err
		}
		return metric, nil
	}
	return nil, nil

}

func GetJson(message Message) ([]byte, error) {
	//TODO wrap in IfDebug
	log.Printf("%v.GetJson()\n", message)
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("json.Marshal for %v failed error: %v\n", message, err)
		return []byte{}, err
	}
	//TODO wrap in IfDebug
	log.Printf("%v.GetJson() returning %s\n", message, string(bytes))
	return bytes, nil
}
