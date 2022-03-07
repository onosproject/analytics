package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/analytics/pkg/messages"
)

func main() {
	//var configFile = flag.String("conf", "generator.json", "json config file for generator")
	var kafkaURI = flag.String("uri", "localhost:9092", "uri of the kafka bus")
	var count = flag.Int("count", 1, "iber of each message")
	flag.Parse()
	log.Println("MessageGenerator")
	//eventWriter := kafkaClient.GetWriter(*kafkaURI, "events")
	go generateAlarms(*kafkaURI, "alarms", *count)
	go generateMetrics(*kafkaURI, "metrics", *count)
	generateEvents(*kafkaURI, "metrics", *count)

}

func generateAlarms(kafkaURI string, topic string, count int) {
	alarmWriter := kafkaClient.GetWriter(kafkaURI, "alarms")
	alarm := messages.Alarm{Type: "TEST", Message: "Generated Alarm"}
	if count > 0 {
		for i := 0; i < count; i++ {
			alarm.Device = "SomeDevice" + strconv.Itoa(i%20)
			alarm.Serverity = uint(i % 4)
			bytes, _ := messages.GetJson(alarm)
			alarmWriter.SendMessage(bytes)
			log.Printf("Send Alarm %s\n", string(bytes))
			time.Sleep(10 * time.Second)
		}
	} else {
		i := 0
		for {
			alarm.Device = "SomeDevice" + strconv.Itoa(i%20)
			alarm.Serverity = uint(i % 4)
			bytes, _ := messages.GetJson(alarm)
			alarmWriter.SendMessage(bytes)
			log.Printf("Send Alarm %s\n", string(bytes))
			time.Sleep(10 * time.Second)
			i++
		}

	}

}

func generateMetrics(kafkaURI string, topic string, count int) {
	metricWriter := kafkaClient.GetWriter(kafkaURI, topic)
	metric := messages.Metric{}
	if count > 0 {
		for i := 0; i < count; i++ {
			metric.Device = "SomeDevice" + strconv.Itoa(i%20)
			metric.BytesIn = uint64(30 * i)
			metric.BytesOut = uint64(30 * i)
			metric.ErrorsIn = uint64(i)
			metric.ErrorsOut = uint64(i)

			bytes, _ := messages.GetJson(metric)
			metricWriter.SendMessage(bytes)
			log.Printf("Send Metric %s\n", string(bytes))
			time.Sleep(10 * time.Second)
		}
	} else {
		for {
			i := 0
			metric.Device = "SomeDevice" + strconv.Itoa(i%20)
			metric.BytesIn = uint64(30 * i)
			metric.BytesOut = uint64(30 * i)
			metric.ErrorsIn = uint64(i)
			metric.ErrorsOut = uint64(i)

			bytes, _ := messages.GetJson(metric)
			metricWriter.SendMessage(bytes)
			time.Sleep(10 * time.Second)
			log.Printf("Send Metric %s\n", string(bytes))
			i++
		}
	}
}
func generateEvents(kafkaURI string, topic string, count int) {
	eventWriter := kafkaClient.GetWriter(kafkaURI, topic)
	event := messages.Event{Type: "UE", Message: "Successful Connect"}
	if count > 0 {
		for i := 0; i < count; i++ {
			event.Device = "SomeDevice" + strconv.Itoa(i%20)

			bytes, _ := messages.GetJson(event)
			eventWriter.SendMessage(bytes)
			log.Printf("Send Event %s\n", string(bytes))
			time.Sleep(10 * time.Second)
		}
	} else {
		for {
			i := 0
			event.Device = "SomeDevice" + strconv.Itoa(i%20)

			bytes, _ := messages.GetJson(event)
			eventWriter.SendMessage(bytes)
			log.Printf("Send Event %s\n", string(bytes))
			time.Sleep(10 * time.Second)
			i++
		}
	}
}
