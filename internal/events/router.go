package events

import (
	"context"
	"encoding/json"
	"fmt"

	kafkago "github.com/segmentio/kafka-go"
)

func RouteEvent(ctx context.Context, msg kafkago.Message) error {
	var event Event

	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}

	switch msg.Topic {

	case TopicPayment:
		return HandlePaymentEvent(ctx, event)

	case TopicOrder:
		return HandleOrderEvent(ctx, event)

	default:
		return fmt.Errorf(
			"unknown topic: %s",
			msg.Topic,
		)
	}
}
