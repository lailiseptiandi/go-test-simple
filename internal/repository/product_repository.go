package repository

import (
	"context"

	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Get(ctx context.Context, req request.ProductListRequest) ([]*models.Product, *utils.Pagination, error)
	Create(ctx context.Context, product models.Product) (*models.Product, error)
	FindByID(ctx context.Context, id int64) (*models.Product, error)
	Update(ctx context.Context, id int64, product models.Product) (*models.Product, error)
	Delete(ctx context.Context, id int64) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Get(ctx context.Context, req request.ProductListRequest) ([]*models.Product, *utils.Pagination, error) {

	var products []*models.Product
	query := r.db.WithContext(ctx).Model(&models.Product{})

	if req.Search != "" {
		search := "%" + req.Search + "%"
		query.Where("name ilike ? OR sku ilike ?", search, search)
	}
	pagination := &utils.Pagination{
		CurrentPage: req.Page,
		PerPage:     int64(req.Limit),
	}

	query.Preload("Category")

	queryPaginate, pagination, err := utils.Paginate(pagination, "products", query)
	if err != nil {
		return nil, nil, err
	}

	err = queryPaginate.Find(&products).Error
	if err != nil {
		return nil, nil, err
	}

	return products, pagination, nil
}

func (r *productRepository) Create(ctx context.Context, product models.Product) (*models.Product, error) {
	err := r.db.WithContext(ctx).FirstOrCreate(&product).Error
	if err != nil {
		return nil, err
	}

	p, err := r.FindByID(ctx, int64(product.ID))
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *productRepository) FindByID(ctx context.Context, id int64) (*models.Product, error) {
	var product *models.Product
	err := r.db.WithContext(ctx).Where("id = ?", id).Preload("Category").Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Update(ctx context.Context, id int64, product models.Product) (*models.Product, error) {
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return nil, err
	}

	p, err := r.FindByID(ctx, int64(product.ID))
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
