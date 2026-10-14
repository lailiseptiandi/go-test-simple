package request

import "mime/multipart"

type ProductListRequest struct {
	Page   int
	Limit  int
	Search string
}

type ProductCreateUpdateRequest struct {
	Name        string                `form:"name" binding:"required"`
	SKU         string                `form:"sku"`
	Image       *multipart.FileHeader `form:"image" `
	Description string                `form:"description"`
	CategoryID  string                `form:"category_id" binding:"required"`
	Price       string                `form:"price"`
	ImagePath   string
}
