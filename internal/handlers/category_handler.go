package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/responses"
	"github.com/lailiseptiandi/go-test-simple/internal/services"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/helpers"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/response"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) Get(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	req := request.CategoryListRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	}

	categories, pagination, err := h.categoryService.Get(c, req)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.CategoryResponse{}).FormatterCategoryResponse(categories)
	response.SuccessWithMessage(c, "Successfully get categories", formatter, pagination)

}

func (h *CategoryHandler) Create(c *gin.Context) {
	name := c.Request.FormValue("name")
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	reqBody := request.CategoryCreateUpdateRequest{
		Name: name,
	}

	if file != nil {
		imagePath := helpers.UploadImageFile(file, header)
		reqBody.Image = imagePath
	}

	category, err := h.categoryService.Create(c, reqBody)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.CategoryResponse{}).FormatterCategoryCreateUpdateResponse(*category)
	response.SuccessWithMessage(c, "Successfully create category", formatter, nil)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}
	category, err := h.categoryService.FindByID(c, id)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	formatter := (&responses.CategoryResponse{}).FormatterCategoryCreateUpdateResponse(*category)
	response.SuccessWithMessage(c, "successfully get category by id", formatter, nil)

}

func (h *CategoryHandler) Update(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}

	name := c.Request.FormValue("name")
	file, header, err := c.Request.FormFile("image")

	reqBody := request.CategoryCreateUpdateRequest{
		Name: name,
		ID:   id,
	}

	if file != nil {
		imagePath := helpers.UploadImageFile(file, header)
		reqBody.Image = imagePath
	}

	category, err := h.categoryService.Update(c, reqBody)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.CategoryResponse{}).FormatterCategoryCreateUpdateResponse(*category)
	response.SuccessWithMessage(c, "Successfully update category", formatter, nil)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}
	err := h.categoryService.Delete(c, id)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.SuccessWithMessage(c, "successfully delete category", nil, nil)
}
