package request

type CategoryListRequest struct {
	Page   int
	Limit  int
	Search string
}

type CategoryCreateUpdateRequest struct {
	ID    string
	Name  string `json:"name" binding:"required"`
	Image string `json:"image"`
}
