package services

import (
	"context"
	"errors"

	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/internal/repository"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, req request.PaymentRequest) error
	CreatePaymentIdempotencyKey(ctx context.Context, idempotencyKey string, req request.PaymentRequest) (*models.Payment, error)
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) *paymentService {
	return &paymentService{paymentRepo: paymentRepo}
}

func (s *paymentService) CreatePayment(ctx context.Context, req request.PaymentRequest) error {

	payment := models.Payment{
		OrderID: req.OrderID,
		Amount:  req.Amount,
	}
	err := s.paymentRepo.CreatePayment(ctx, payment)
	if err != nil {
		return err
	}

	return nil
}

func (s *paymentService) CreatePaymentIdempotencyKey(ctx context.Context, idempotencyKey string, req request.PaymentRequest) (*models.Payment, error) {

	payment := &models.Payment{
		OrderID:        req.OrderID,
		IdempotencyKey: idempotencyKey,
		Amount:         req.Amount,
		Status:         "SUCCESS",
	}
	created, err := s.paymentRepo.CreatePaymentIdempotencyCase(ctx, *payment)
	if err != nil {
		return nil, err
	}

	if created {
		return payment, nil
	}

	// existing idempotencyKey
	existing, err := s.paymentRepo.FindByIdempotencyKey(ctx, idempotencyKey)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		return nil, errors.New("payment not found after conflict")
	}

	if existing.OrderID != req.OrderID {
		return nil, errors.New("idempotency key already used with different payload")
	}

	return existing, nil
}
