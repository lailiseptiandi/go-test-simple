package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/internal/repository"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/jwt"
)

type AuthService interface {
	Login(ctx context.Context, req request.UserLoginRequest) (*models.User, error)
	Register(ctx context.Context, req request.UserRegisterRequest) (*models.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *authService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(ctx context.Context, req request.UserLoginRequest) (*models.User, error) {

	user, err := s.userRepo.FindByEmailOrUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	// compare password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	// generate access token & refresh token
	token, err := jwt.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	user.AccessToken = token
	user.RefreshToken = refreshToken

	// update token
	go func() {
		_, err = s.userRepo.Update(ctx, int64(user.ID), user)
		if err != nil {
			fmt.Printf("err: %v", err.Error())
		}
	}()

	return user, nil
}

func (s *authService) Register(ctx context.Context, req request.UserRegisterRequest) (*models.User, error) {

	// validate password == confirm
	if req.Password != req.PasswordConfirmation {
		return nil, errors.New("password confirmation not match password")
	}

	// validate password strenth
	err := validatePasswordStrength(req.Password)
	if err != nil {
		return nil, err
	}

	data := models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Roles:    models.ROLE_USER, // default user
		Email:    req.Email,
	}
	user, err := s.userRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	// generate access token & refresh token
	token, err := jwt.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	user.AccessToken = token
	user.RefreshToken = refreshToken

	// update token
	go func() {
		_, err = s.userRepo.Update(ctx, int64(user.ID), user)
		if err != nil {
			fmt.Printf("err: %v", err.Error())
		}
	}()

	return user, nil
}

func validatePasswordStrength(password string) error {

	if len(password) < 8 {
		return errors.New("password must be at least 8 character long")
	}

	if !hasNumber(password) {
		return errors.New("password must containt at least one number")
	}

	if !hasSymbol(password) {
		return errors.New("password must containt at least one symbol")
	}

	if !hasUpperCase(password) {
		return errors.New("password must containt at least one uppercase latter")
	}

	if !hasLowerCase(password) {
		return errors.New("password must containt at least one lowercase latter")
	}

	return nil
}

func hasNumber(password string) bool {
	for _, c := range password {
		if c > '0' && c <= '9' {
			return true
		}
	}

	return false
}

func hasSymbol(password string) bool {
	for _, c := range password {
		if (c >= '!' && c <= '/') || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= '~') {
			return true
		}
	}

	return false
}

func hasUpperCase(password string) bool {
	for _, c := range password {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}

	return false
}

func hasLowerCase(password string) bool {
	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}

	return false
}
