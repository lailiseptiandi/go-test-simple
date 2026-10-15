package kafka

import (
	"context"
	"encoding/json"

	kafkago "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafkago.Writer
}

func NewProducer(broker string) *Producer {
	return &Producer{
		writer: &kafkago.Writer{
			Addr:     kafkago.TCP(broker),
			Balancer: &kafkago.LeastBytes{},
		},
	}
}

func (p *Producer) Publish(ctx context.Context, topic string, key string, data any) error {

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, kafkago.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: payload,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
