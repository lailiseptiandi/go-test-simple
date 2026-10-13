package utils

import (
	"fmt"
	"math"
	"os"

	"gorm.io/gorm"
)

type Pagination struct {
	CurrentPage  int     `json:"current_page"`
	NextPage     *string `json:"next_page"`
	PreviousPage *string `json:"previous_page"`
	TotalPages   int64   `json:"total_pages"`
	PerPage      int64   `json:"per_page"`
	LastPage     int64   `json:"last_page"`
	Total        int64   `json:"total"`
}

func Paginate(pagination *Pagination, endpoint string, query *gorm.DB) (*gorm.DB, *Pagination, error) {
	limit := pagination.GetLimit()
	page := pagination.GetPage()

	var totalRows int64
	if err := query.Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}
	pagination.Total = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))
	pagination.TotalPages = int64(totalPages)
	pagination.LastPage = int64(totalPages)

	queryResult := query.Offset(pagination.GetOffset()).Limit(limit)

	appURL := os.Getenv("APP_URL")

	if page < totalPages {
		nextPage := fmt.Sprintf("%s/%s?page=%d&limit=%d", appURL, endpoint, page+1, limit)
		pagination.NextPage = &nextPage
	}

	if page > 1 && totalPages > 1 {
		prevPage := fmt.Sprintf("%s/%s?page=%d&limit=%d", appURL, endpoint, page-1, limit)
		pagination.PreviousPage = &prevPage
	}

	return queryResult, pagination, nil
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.PerPage <= 0 {
		p.PerPage = 10
	}
	return int(p.PerPage)
}

func (p *Pagination) GetPage() int {
	if p.CurrentPage <= 0 {
		p.CurrentPage = 1
	}
	return int(p.CurrentPage)
}
