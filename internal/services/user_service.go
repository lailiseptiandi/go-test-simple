package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/internal/repository"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/helpers"
)

type UserService interface {
	Create(ctx context.Context, req request.UserRequest) (*models.User, error)
	Get(req request.UserListRequest) ([]*models.User, error)
	FindByID(id string) (*models.User, error)
	Update(ctx context.Context, req request.UserRequest) (*models.User, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Create(ctx context.Context, req request.UserRequest) (*models.User, error) {

	hashPassword, _ := utils.HashPassword(req.Password)
	userData := models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: hashPassword,
		Email:    req.Email,
		Roles:    1, // harcode user role
	}

	user, err := s.userRepo.Create(ctx, userData)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Get(req request.UserListRequest) ([]*models.User, error) {
	users, err := s.userRepo.Get(req)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) FindByID(id string) (*models.User, error) {

	decryptID, _ := helpers.Decrypt(id)
	user, err := s.userRepo.FindByID(int64(decryptID))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(ctx context.Context, req request.UserRequest) (*models.User, error) {

	decryptID, err := helpers.Decrypt(req.ID)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return nil, errors.New("id not found")
	}

	// validate is existing data
	userDetail, err := s.userRepo.FindByID(int64(decryptID))
	if err != nil {
		return nil, err
	}

	if userDetail.ID == 0 {
		return nil, errors.New("user not found")
	}

	data := &models.User{
		Password: req.Password,
		Name:     req.Name,
		Roles:    1, // harcode super admin
	}

	// skip unique data
	if req.Username != userDetail.Username {
		data.Username = req.Username
	}

	if req.Email != userDetail.Email {
		data.Email = req.Email
	}

	if req.Password != "" {
		// check password
		hashPassword, _ := utils.HashPassword(req.Password)
		isMatchPassword := utils.CheckPasswordHash(hashPassword, userDetail.Password)

		// new value if not match password old
		if !isMatchPassword {
			data.Password = hashPassword
		}
	}

	user, err := s.userRepo.Update(ctx, int64(decryptID), data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	decryptID, err := helpers.Decrypt(id)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return errors.New("id not found")
	}

	err = s.userRepo.Delete(ctx, int64(decryptID))
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) ValidateUser(user models.User) error {
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	existingUser, err := s.userRepo.FindByEmail(user.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	return nil

}
