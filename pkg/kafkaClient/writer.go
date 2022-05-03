package kafkaClient

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Writer struct {
	kafkaWriter kafka.Writer
}

/*
GetWriter
creates a kafka.Writer and wraps in a Writer stucture
*/
func GetWriter(kafkaURI string, topic string) Writer {
	log.Debugf("GetWriter(%s,%s)", kafkaURI, topic)
	producer := kafka.Writer{
		Addr:     kafka.TCP(kafkaURI),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	writer := Writer{
		kafkaWriter: producer,
	}
	return writer
}

/*
SendMessage
constructs a kafkaMessage from message and writes to
the topic the writer is attached to
*/
func (writer Writer) SendMessage(message []byte) error {
	log.Debugf("SendMessage(%s)", string(message))
	msg := kafka.Message{Value: message}
	err := writer.kafkaWriter.WriteMessages(context.Background(), msg)
	return err
}
