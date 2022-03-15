package processor

import (
	"strings"

	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/analytics/pkg/logger"
)

type ProcessorType uint

const (
	ALARM ProcessorType = iota
	EVENT
	METRIC
)

func StartProcessor(channelName string, messageChan chan string, errorChan chan error, kafkaURI []string, kafkaTopic string) {

	writer := kafkaClient.GetWriter(kafkaURI[0], kafkaTopic)
	var myType ProcessorType
	logger.Debug(strings.ToUpper(channelName))
	switch strings.ToUpper(channelName) {
	case "METRICS":
		logger.Debug("metric")
		myType = METRIC
	case "EVENTS":
		logger.Debug("event")
		myType = EVENT
	case "ALARMS":
		logger.Debug("alarm")
		myType = ALARM

	}

	logger.Debug("StartProcessor(%s,%v,%s) type: %d", channelName, kafkaURI, kafkaTopic, myType)

	for {
		select {
		case err := <-errorChan:
			logger.Error("kafkaClient read error: %v", err)
		case messageJSON := <-messageChan:
			if logger.IfDebug() {
				logger.Debug("MessageProcessor received %s", messageJSON)
			}
			switch myType {
			case ALARM:
				message, err := processAlarm(messageJSON)
				if err != nil {
					logger.Error("Failed to process Alarm: %s err:%v", messageJSON, err)
				}

				err = writer.SendMessage(message)
				if err != nil {
					logger.Error("Failed to send message: %s", message)
				}
			case EVENT:
				message, err := processEvent(messageJSON)
				if err != nil {
					logger.Error("Failed to process Event: %s err:%v", messageJSON, err)
				}

				err = writer.SendMessage(message)
				if err != nil {
					logger.Error("Failed to send message: %s", message)
				}
			case METRIC:
				message, err := processMetric(messageJSON)
				if err != nil {
					logger.Error("Failed to process Metric: %s err:%v", messageJSON, err)
				}

				err = writer.SendMessage(message)
				if err != nil {
					logger.Error("Failed to send message: %s", message)
				}
			}
		}
	}
}
