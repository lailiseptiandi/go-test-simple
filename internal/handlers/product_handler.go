package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/responses"
	"github.com/lailiseptiandi/go-test-simple/internal/services"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/response"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) Get(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	req := request.ProductListRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	}

	products, pagination, err := h.productService.Get(c, req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.ProductResponse{}).FormatterGetProductResponse(products)
	response.SuccessWithMessage(c, "Successfully get products", formatter, pagination)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req request.ProductCreateUpdateRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Println("error: ", req)
		response.BadRequest(c, err.Error())
		return
	}

	if req.Image != nil {

		storageDir := os.Getenv("STORAGE_PATH")
		if storageDir == "" {
			storageDir = "storage"
		}
		if err := os.MkdirAll(storageDir, 0750); err != nil {
			response.InternalServerError(c, err.Error())
			return
		}
		path := filepath.Join(storageDir, req.Image.Filename)
		if err := c.SaveUploadedFile(req.Image, path); err != nil {
			response.InternalServerError(c, err.Error())
			return
		}

		req.ImagePath = "/" + path

	}

	product, err := h.productService.Create(c, req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.ProductResponse{}).FormatterCreateUpdateProductResponse(product)
	response.SuccessWithMessage(c, "successfully create products ", formatter, nil)
}
func (h *ProductHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}
	product, err := h.productService.FindByID(c, id)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	formatter := (&responses.ProductResponse{}).FormatterCreateUpdateProductResponse(product)
	response.SuccessWithMessage(c, "successfully get product by id", formatter, nil)
}
func (h *ProductHandler) Update(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}

	var req request.ProductCreateUpdateRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Println("error: ", req)
		response.BadRequest(c, err.Error())
		return
	}
	if req.Image != nil {
		path := fmt.Sprintf("/%v/%v", os.Getenv("STORAGE_PATH"), req.Image.Filename)
		if err := c.SaveUploadedFile(req.Image, path); err != nil {
			response.InternalServerError(c, err.Error())
			return
		}

		req.ImagePath = path

	}

	product, err := h.productService.Update(c, id, req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.ProductResponse{}).FormatterCreateUpdateProductResponse(product)
	response.SuccessWithMessage(c, "successfully update products ", formatter, nil)
}
func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}
	err := h.productService.Delete(c, id)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "successfully delete product", nil, nil)
}
