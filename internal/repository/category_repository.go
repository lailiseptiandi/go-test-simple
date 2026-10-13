package repository

import (
	"context"

	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Get(ctx context.Context, req request.CategoryListRequest) ([]*models.Category, *utils.Pagination, error)
	Create(ctx context.Context, category models.Category) (*models.Category, error)
	FindByID(ctx context.Context, ID uint) (*models.Category, error)
	Update(ctx context.Context, ID uint, category models.Category) (*models.Category, error)
	Delete(ctx context.Context, id int64) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Get(ctx context.Context, req request.CategoryListRequest) ([]*models.Category, *utils.Pagination, error) {
	var categories []*models.Category
	query := r.db.Model(&models.Category{}).WithContext(ctx)
	if req.Search != "" {
		search := "%" + req.Search + "%"
		query.Where("name ilike ?", search)
	}

	pagination := &utils.Pagination{
		CurrentPage: req.Page,
		PerPage:     int64(req.Page),
	}

	queryPaginate, pagination, err := utils.Paginate(pagination, "category", query)
	if err != nil {
		return nil, nil, err
	}

	err = queryPaginate.Find(&categories).Error
	if err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil

}
func (r *categoryRepository) Create(ctx context.Context, category models.Category) (*models.Category, error) {

	err := r.db.Model(&models.Category{}).FirstOrCreate(&category).Error
	if err != nil {
		return nil, err
	}

	c, err := r.FindByID(ctx, category.ID)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *categoryRepository) FindByID(ctx context.Context, ID uint) (*models.Category, error) {
	var category *models.Category
	err := r.db.Model(&models.Category{}).WithContext(ctx).Where("id = ?", ID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Update(ctx context.Context, ID uint, category models.Category) (*models.Category, error) {
	err := r.db.Model(&models.Category{}).Where("id = ? ", ID).Updates(&category).Error
	if err != nil {
		return nil, err
	}

	c, err := r.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *categoryRepository) Delete(ctx context.Context, id int64) error {

	err := r.db.WithContext(ctx).Model(&models.Category{}).Where("id = ?", id).Delete(&models.Category{}).Error
	if err != nil {
		return err
	}
	return nil
}
