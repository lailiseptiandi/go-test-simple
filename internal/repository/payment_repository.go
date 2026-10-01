package repository

import (
	"context"
	"errors"

	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, req models.Payment) error
	CreatePaymentIdempotencyCase(ctx context.Context, payment models.Payment) (bool, error)
	FindByIdempotencyKey(ctx context.Context, idempotencyKey string) (*models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) paymentRepository {
	return paymentRepository{db: db}
}

func (r *paymentRepository) CreatePayment(ctx context.Context, payment models.Payment) error {
	err := r.db.WithContext(ctx).Create(&payment).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) CreatePaymentIdempotencyCase(ctx context.Context, payment models.Payment) (bool, error) {

	result := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "idempotency_key"},
			},
			DoNothing: true,
		}).Create(&payment)

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected == 1, nil
}

func (r *paymentRepository) FindByIdempotencyKey(ctx context.Context, idempotencyKey string) (*models.Payment, error) {
	var payment *models.Payment

	err := r.db.WithContext(ctx).Where("idempotency_key = ?", idempotencyKey).First(&payment).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return payment, nil
}
