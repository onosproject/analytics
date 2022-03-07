package kafkaClient

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Writer struct {
	kafkaWriter kafka.Writer
}

func GetWriter(kafkaURI string, topic string) Writer {
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
func (writer Writer) SendMessage(message []byte) error {
	msg := kafka.Message{Value: message}
	err := writer.kafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	}
	return nil
}
