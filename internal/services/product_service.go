package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/internal/repository"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/helpers"
)

type ProductService interface {
	Get(ctx context.Context, req request.ProductListRequest) ([]*models.Product, *utils.Pagination, error)
	Create(ctx context.Context, req request.ProductCreateUpdateRequest) (*models.Product, error)
	FindByID(ctx context.Context, id string) (*models.Product, error)
	Update(ctx context.Context, id string, req request.ProductCreateUpdateRequest) (*models.Product, error)
	Delete(ctx context.Context, id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService {
	return &productService{repo: repo}
}

func (s *productService) Get(ctx context.Context, req request.ProductListRequest) ([]*models.Product, *utils.Pagination, error) {
	products, pagination, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return products, pagination, nil
}
func (s *productService) Create(ctx context.Context, req request.ProductCreateUpdateRequest) (*models.Product, error) {

	convertPrice, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return nil, errors.New("error convert price product")
	}

	categoryID, _ := helpers.Decrypt(req.CategoryID)

	product := models.Product{
		Name:        req.Name,
		Slug:        slug.MakeLang(req.Name, "id"),
		SKU:         req.SKU,
		Image:       req.ImagePath,
		Description: req.Description,
		CategoryID:  int64(categoryID),
		Price:       convertPrice,
	}

	res, err := s.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *productService) FindByID(ctx context.Context, id string) (*models.Product, error) {

	productID, _ := helpers.Decrypt(id)
	product, err := s.repo.FindByID(ctx, int64(productID))
	if err != nil {
		return nil, err
	}
	return product, nil

}
func (s *productService) Update(ctx context.Context, id string, req request.ProductCreateUpdateRequest) (*models.Product, error) {

	convertPrice, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return nil, errors.New("error convert price product")
	}
	categoryID, _ := helpers.Decrypt(req.CategoryID)
	product := models.Product{
		Name:        req.Name,
		Slug:        slug.MakeLang(req.Name, "id"),
		SKU:         req.SKU,
		Image:       req.ImagePath,
		Description: req.Description,
		CategoryID:  int64(categoryID),
		Price:       convertPrice,
	}

	productID, _ := helpers.Decrypt(id)

	res, err := s.repo.Update(ctx, int64(productID), product)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *productService) Delete(ctx context.Context, id string) error {
	productID, _ := helpers.Decrypt(id)
	err := s.repo.Delete(ctx, int64(productID))
	if err != nil {
		return err
	}
	return nil
}
