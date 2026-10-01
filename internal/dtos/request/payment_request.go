package request

type PaymentRequest struct {
	OrderID string  `json:"order_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required,gt=0"`
}
