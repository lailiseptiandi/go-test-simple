package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/responses"
	"github.com/lailiseptiandi/go-test-simple/internal/services"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/response"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Create(c *gin.Context) {
	var reqBody request.UserRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ValidationError(c, reqBody)
		return
	}

	user, err := h.userService.Create(c, reqBody)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.UserResponse{}).FormatterUserResponse(*user)
	response.Created(c, "successfully create user", formatter)
}

func (h *UserHandler) Get(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 20
	}
	req := request.UserListRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	}

	users, err := h.userService.Get(req)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.UserResponse{}).FormatterGetUserResponse(users)
	response.SuccessWithMessage(c, "Successfully get all users", formatter)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}

	user, err := h.userService.FindByID(id)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.UserResponse{}).FormatterUserResponse(*user)
	response.SuccessWithMessage(c, "Successfully get user by id", formatter)
}

func (h *UserHandler) Update(c *gin.Context) {
	var reqBody request.UserRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ValidationError(c, reqBody)
		return
	}

	// validate param
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}

	reqBody.ID = id

	user, err := h.userService.Update(c, reqBody)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.UserResponse{}).FormatterUserResponse(*user)
	response.Created(c, "successfully create user", formatter)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.Error(c, http.StatusBadRequest, "id param is required")
		return
	}

	err := h.userService.Delete(c, id)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.SuccessWithMessage(c, "Successfully deleted user", nil)
}
