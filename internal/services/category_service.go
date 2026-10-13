package services

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/internal/repository"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/helpers"
)

type CategoryService interface {
	Get(ctx context.Context, req request.CategoryListRequest) ([]*models.Category, *utils.Pagination, error)
	Create(ctx context.Context, req request.CategoryCreateUpdateRequest) (*models.Category, error)
	FindByID(ctx context.Context, categoryID string) (*models.Category, error)
	Update(ctx context.Context, req request.CategoryCreateUpdateRequest) (*models.Category, error)
	Delete(ctx context.Context, id string) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *categoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Get(ctx context.Context, req request.CategoryListRequest) ([]*models.Category, *utils.Pagination, error) {

	categories, pagination, err := s.repo.Get(ctx, req)
	if err != nil {
		return categories, pagination, err
	}

	return categories, pagination, nil
}

func (s *categoryService) Create(ctx context.Context, req request.CategoryCreateUpdateRequest) (*models.Category, error) {
	data := models.Category{
		Name:  req.Name,
		Slug:  slug.MakeLang(req.Name, "id"),
		Image: req.Image,
	}
	category, err := s.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) FindByID(ctx context.Context, categoryID string) (*models.Category, error) {

	id, _ := helpers.Decrypt(categoryID)
	category, err := s.repo.FindByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) Update(ctx context.Context, req request.CategoryCreateUpdateRequest) (*models.Category, error) {
	data := models.Category{
		Name:  req.Name,
		Slug:  slug.MakeLang(req.Name, "id"),
		Image: req.Image,
	}

	id, _ := helpers.Decrypt(req.ID)

	category, err := s.repo.Update(ctx, uint(id), data)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) Delete(ctx context.Context, categoryID string) error {
	id, _ := helpers.Decrypt(categoryID)
	err := s.repo.Delete(ctx, int64(id))
	if err != nil {
		return err
	}

	return nil
}
