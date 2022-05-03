package processor

import (
	"strings"

	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

type ProcessorType uint

const (
	ALARM ProcessorType = iota
	EVENT
	METRIC
)

var log = logging.GetLogger("message_processor")

func StartProcessor(channelName string, messageChan chan string, errorChan chan error, kafkaURI []string, kafkaTopic string) {

	writer := kafkaClient.GetWriter(kafkaURI[0], kafkaTopic)
	var myType ProcessorType
	log.Debug(strings.ToUpper(channelName))
	switch strings.ToUpper(channelName) {
	case "METRICS":
		log.Debug("metric")
		myType = METRIC
	case "EVENTS":
		log.Debug("event")
		myType = EVENT
	case "ALARMS":
		log.Debug("alarm")
		myType = ALARM

	}

	log.Debugf("StartProcessor(%s,%v,%s) type: %d", channelName, kafkaURI, kafkaTopic, myType)

	for {
		select {
		case err := <-errorChan:
			log.Errorf("kafkaClient read error: %v", err)
		case messageJSON := <-messageChan:
			log.Debugf("MessageProcessor received %s", messageJSON)
			switch myType {
			case ALARM:
				message, err := processAlarm(messageJSON)
				if err != nil {
					log.Errorf("Failed to process Alarm: %s err:%v", messageJSON, err)
				}

				err = writer.SendMessage(message)
				if err != nil {
					log.Errorf("Failed to send message: %s", message)
				}
			case EVENT:
				message, err := processEvent(messageJSON)
				if err != nil {
					log.Errorf("Failed to process Event: %s err:%v", messageJSON, err)
				}

				err = writer.SendMessage(message)
				if err != nil {
					log.Errorf("Failed to send message: %s", message)
				}
			case METRIC:
				message, err := processMetric(messageJSON)
				if err != nil {
					log.Errorf("Failed to process Metric: %s err:%v", messageJSON, err)
				}

				err = writer.SendMessage(message)
				if err != nil {
					log.Errorf("Failed to send message: %s", message)
				}
			}
		}
	}
}
