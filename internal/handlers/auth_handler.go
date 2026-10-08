package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/responses"
	"github.com/lailiseptiandi/go-test-simple/internal/services"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/response"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.UserLoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.authService.Login(c, req)

	if err == gorm.ErrRecordNotFound {
		response.NotFound(c, err.Error())
		return
	}
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.UserResponse{}).FormatterLoginResponse(*user)
	response.SuccessWithMessage(c, "successfully login", formatter)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req request.UserRegisterRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.authService.Register(c, req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	formatter := (&responses.UserResponse{}).FormatterLoginResponse(*user)
	response.SuccessWithMessage(c, "successfully register user", formatter)
}
