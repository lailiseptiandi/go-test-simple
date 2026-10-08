package repository

import (
	"context"

	"github.com/lailiseptiandi/go-test-simple/internal/dtos/request"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (*models.User, error)
	Get(req request.UserListRequest) ([]*models.User, error)
	FindByID(id int64) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByEmailOrUsername(ctx context.Context, emailUsername string) (*models.User, error)
	Update(ctx context.Context, id int64, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user models.User) (*models.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Get(req request.UserListRequest) ([]*models.User, error) {
	var users []*models.User
	query := r.db

	page := req.Page
	pageSize := req.Limit

	if req.Search != "" {
		search := "%" + req.Search + "%"
		query = query.Where("name ilike ? OR username ilike ?", search, search)
	}

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByID(id int64) (*models.User, error) {
	var user *models.User
	err := r.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmailOrUsername(ctx context.Context, emailUsername string) (*models.User, error) {
	var user *models.User
	err := r.db.WithContext(ctx).Where("email = ? OR username = ?", emailUsername, emailUsername).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepository) Update(ctx context.Context, id int64, user *models.User) (*models.User, error) {

	err := r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	var userData *models.User
	err = r.db.Where("id = ?", id).First(&userData).Error
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
