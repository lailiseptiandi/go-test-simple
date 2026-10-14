package responses

import (
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/pkg/config"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/helpers"
)

type ProductResponse struct {
	ID          string           ` json:"id"`
	Name        string           `json:"name"`
	SKU         string           `json:"sku"`
	Slug        string           `json:"slug"`
	Image       string           `json:"image"`
	Description string           `json:"description"`
	CategoryID  string           `json:"category_id"`
	Category    CategoryResponse `gorm:"foreignKey:ID;references:CategoryID"`
	Price       float64          `json:"price"`
}

func (res *ProductResponse) FormatterGetProductResponse(products []*models.Product) []*ProductResponse {
	var data []*ProductResponse
	config, _ := config.LoadConfig()
	for _, product := range products {
		encryptID, _ := helpers.Encrypt(uint64(product.ID))
		categoryID, _ := helpers.Encrypt(uint64(product.CategoryID))
		productResult := ProductResponse{
			ID:          encryptID,
			Name:        product.Name,
			SKU:         product.SKU,
			Slug:        product.Slug,
			Image:       config.AppURL + product.Image,
			Description: product.Description,
			CategoryID:  categoryID,
			Category: CategoryResponse{
				ID:    categoryID,
				Name:  product.Category.Name,
				Slug:  product.Category.Slug,
				Image: config.AppURL + product.Category.Image,
			},
			Price: product.Price,
		}
		data = append(data, &productResult)
	}
	return data
}

func (res *ProductResponse) FormatterCreateUpdateProductResponse(product *models.Product) *ProductResponse {
	config, _ := config.LoadConfig()
	encryptID, _ := helpers.Encrypt(uint64(product.ID))
	categoryID, _ := helpers.Encrypt(uint64(product.CategoryID))
	productResult := &ProductResponse{
		ID:          encryptID,
		Name:        product.Name,
		SKU:         product.SKU,
		Slug:        product.Slug,
		Image:       config.AppURL + product.Image,
		Description: product.Description,
		CategoryID:  categoryID,
		Category: CategoryResponse{
			ID:    categoryID,
			Name:  product.Category.Name,
			Slug:  product.Category.Slug,
			Image: config.AppURL + product.Category.Image,
		},
		Price: product.Price,
	}
	return productResult
}
