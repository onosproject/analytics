package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/analytics/pkg/messages"
)

var metricsSent = 0
var alarmsSent = 0
var eventsSent = 0
var refCount = 0

func main() {
	//var configFile = flag.String("conf", "generator.json", "json config file for generator")
	var kafkaURI = flag.String("uri", "localhost:9092", "uri of the kafka bus")
	var count = flag.Int("count", 1, "number of each message to generate")
	var genAlarm = flag.Bool("alarm", false, "generate alarm messages")
	var genEvents = flag.Bool("event", false, "generate event messages")
	var genMetrics = flag.Bool("metrics", false, "generate metric messages")
	var allMetrics = flag.Bool("all", false, "generate all message type")
	var debug = flag.Bool("debug", false, "log debug messages")
	var interval = flag.Int("interval", 10, "number of seconds to sleep between each message type")
	var logFile = flag.String("logFile", "", "file to log to")
	flag.Parse()
	if *logFile != "" {

		file, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("Unable to open %s for writing", *logFile)
			log.Printf("Will print log messages to stdout")
		} else {
			log.SetOutput(file)
		}

	}

	log.Println("MessageGenerator")

	if *allMetrics || *genAlarm {
		refCount++
		go generateAlarms(*kafkaURI, "alarms", *count, *interval, *debug)
	}
	if *allMetrics || *genEvents {
		refCount++
		go generateEvents(*kafkaURI, "events", *count, *interval, *debug)
	}
	if *allMetrics || *genMetrics {
		refCount++
		go generateMetrics(*kafkaURI, "metrics", *count, *interval, *debug)
	}

	statDuration := time.Duration(*interval*2) * time.Second
	for {
		log.Printf("Message Generator stats: Events Sent: %d, Alarms Sent: %d, Metrics Sent: %d\n", eventsSent, alarmsSent, metricsSent)
		if refCount == 0 {
			return
		}
		time.Sleep(statDuration)
	}

}

func generateAlarms(kafkaURI string, topic string, count int, interval int, debug bool) {
	alarmWriter := kafkaClient.GetWriter(kafkaURI, "alarms")
	alarm := messages.Alarm{Type: "TEST", Message: "Generated Alarm"}
	sleepTime := time.Duration(interval) * time.Second
	for {
		alarm.Device = "SomeDevice" + strconv.Itoa(alarmsSent%20)
		alarm.Serverity = uint(alarmsSent % 4)
		bytes, _ := messages.GetJson(alarm)
		alarmWriter.SendMessage(bytes)
		alarmsSent++

		if debug {
			log.Printf("Send Alarm %s\n", string(bytes))
		}
		time.Sleep(sleepTime)
		if count < 1 {
			break
		}
		if alarmsSent >= count {
			refCount--
			return
		}
	}
}

func generateMetrics(kafkaURI string, topic string, count int, interval int, debug bool) {
	sleepTime := time.Duration(interval) * time.Second
	metricWriter := kafkaClient.GetWriter(kafkaURI, topic)
	metric := messages.Metric{}
	for {
		metric.Device = "SomeDevice" + strconv.Itoa(metricsSent%20)
		metric.BytesIn = uint64(30 * metricsSent)
		metric.BytesOut = uint64(30 * metricsSent)
		metric.ErrorsIn = uint64(metricsSent)
		metric.ErrorsOut = uint64(metricsSent)

		bytes, _ := messages.GetJson(metric)
		metricWriter.SendMessage(bytes)
		if debug {
			log.Printf("Send Metric %s\n", string(bytes))
		}
		metricsSent++
		time.Sleep(sleepTime)
		if count < 1 {
			break
		}
		if metricsSent >= count {
			refCount--
			return
		}
	}
}
func generateEvents(kafkaURI string, topic string, count int, interval int, debug bool) {
	sleepTime := time.Duration(interval) * time.Second
	eventWriter := kafkaClient.GetWriter(kafkaURI, topic)
	event := messages.Event{Type: "UE", Message: "Successful Connect"}
	for {
		event.Device = "SomeDevice" + strconv.Itoa(eventsSent%20)

		bytes, _ := messages.GetJson(event)
		eventWriter.SendMessage(bytes)
		if debug {
			log.Printf("Send Event %s\n", string(bytes))
		}
		eventsSent++
		time.Sleep(sleepTime)
		if count < 1 {
			break
		}
		if eventsSent >= count {
			refCount--
			return
		}
	}
}
