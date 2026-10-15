package events

import (
	"context"
	"encoding/json"
	"fmt"
)

type OrderData struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

func HandleOrderEvent(ctx context.Context, event Event) error {
	var order OrderData

	if err := json.Unmarshal(event.Data, &order); err != nil {
		return err
	}

	switch event.EventType {

	case OrderCreated:
		fmt.Println("Order dibuat:", order.OrderID)

	case OrderCancelled:
		fmt.Println("Order dibatalkan:", order.OrderID)

	default:
		return fmt.Errorf(
			"unknown order event: %s",
			event.EventType,
		)
	}

	return nil
}
