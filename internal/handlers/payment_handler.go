package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/services"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/response"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var reqBody request.PaymentRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ValidationError(c, reqBody)
		return
	}

	err := h.paymentService.CreatePayment(c, reqBody)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Created(c, "Successfully created payment", nil)
}

func (h *PaymentHandler) CreatePaymentIdempotencyKey(c *gin.Context) {
	idempotencyKey := c.GetHeader("Idempotency-Key")

	if idempotencyKey == "" {
		response.Error(c, http.StatusBadRequest, "Idempotency-Key header is required")
		return
	}

	var req request.PaymentRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.ValidationError(c, req)
		return
	}

	payment, err := h.paymentService.CreatePaymentIdempotencyKey(c, idempotencyKey, req)
	if err != nil {
		errMsg := errors.New("idempotency key already used with different payload")
		if errors.Is(err, errMsg) {
			response.Error(c, http.StatusConflict, err.Error())
		}
		response.InternalServerError(c, err.Error())
		return
	}

	response.Created(c, "successfully create payment", payment)
}
