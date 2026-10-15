package kafka

import (
	"context"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

type MessageHandler func(ctx context.Context, msg kafkago.Message) error

type Consumer struct {
	reader *kafkago.Reader
}

func NewConsumer(broker string, topic string, groupID string) *Consumer {
	return &Consumer{reader: kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})}
}

func (c *Consumer) Start(ctx context.Context, handler MessageHandler) error {
	defer c.reader.Close()

	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return err
		}

		log.Printf(
			"[KAFKA] Received topic=%s partition=%d offset=%d",
			msg.Topic,
			msg.Partition,
			msg.Offset,
		)

		// businer logic
		if err := handler(ctx, msg); err != nil {
			return err
		}

		// commit if successfully
		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			return err
		}

		log.Printf(
			"[KAFKA] Committed topic=%s offset=%d",
			msg.Topic,
			msg.Offset,
		)

	}
}
