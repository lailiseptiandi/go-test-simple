package responses

import (
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/pkg/config"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/helpers"
)

type CategoryResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}

func (res *CategoryResponse) FormatterCategoryResponse(categories []*models.Category) []*CategoryResponse {
	var results []*CategoryResponse
	config, _ := config.LoadConfig()

	for _, category := range categories {
		encryptID, _ := helpers.Encrypt(uint64(category.ID))
		result := &CategoryResponse{
			ID:    encryptID,
			Name:  category.Name,
			Slug:  category.Slug,
			Image: config.AppURL + category.Image,
		}

		results = append(results, result)
	}

	return results

}

func (res *CategoryResponse) FormatterCategoryCreateUpdateResponse(category models.Category) *CategoryResponse {

	config, _ := config.LoadConfig()
	encryptID, _ := helpers.Encrypt(uint64(category.ID))
	return &CategoryResponse{
		ID:    encryptID,
		Name:  category.Name,
		Slug:  category.Slug,
		Image: config.AppURL + category.Image,
	}
}
