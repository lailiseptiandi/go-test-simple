package events

const (
	TopicPayment = "payment.events"
	TopicOrder   = "order.events"
	TopicUser    = "user.events"
)

const (
	PaymentCreated = "payment.created"
	PaymentSuccess = "payment.success"
	PaymentFailed  = "payment.failed"

	OrderCreated   = "order.created"
	OrderCancelled = "order.cancelled"

	UserRegistered = "user.registered"
)
