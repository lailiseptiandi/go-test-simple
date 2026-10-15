package events

import (
	"context"
	"encoding/json"
	"fmt"
)

type PaymentData struct {
	PaymentID uint   `json:"payment_id"`
	OrderID   string `json:"order_id"`
	Status    string `json:"status"`
}

func HandlePaymentEvent(ctx context.Context, event Event) error {
	var payment PaymentData

	if err := json.Unmarshal(event.Data, &payment); err != nil {
		return err
	}

	switch event.EventType {

	case PaymentCreated:
		fmt.Println("Payment dibuat:", payment.OrderID)

	case PaymentSuccess:
		fmt.Println("Payment berhasil:", payment.OrderID)

	case PaymentFailed:
		fmt.Println("Payment gagal:", payment.OrderID)

	default:
		return fmt.Errorf(
			"unknown payment event: %s",
			event.EventType,
		)
	}

	return nil
}
