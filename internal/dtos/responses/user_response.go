package responses

import (
	"time"

	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/helpers"
)

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func ConvertRoleName(role int) string {
	switch role {
	case 1:
		return "Super Admin"
	case 2:
		return "Pelanggan"
	default:
		return ""
	}
}

func (us *UserResponse) FormatterUserResponse(user models.User) *UserResponse {

	encryptID, _ := helpers.Encrypt(uint64(user.ID))

	return &UserResponse{
		ID:        encryptID,
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		Role:      ConvertRoleName(int(user.Roles)),
		CreatedAt: user.CreatedAt,
	}
}

func (us *UserResponse) FormatterGetUserResponse(users []*models.User) []*UserResponse {
	var datas []*UserResponse

	for _, user := range users {
		encryptID, _ := helpers.Encrypt(uint64(user.ID))

		data := UserResponse{
			ID:        encryptID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      ConvertRoleName(int(user.Roles)),
			CreatedAt: user.CreatedAt,
		}

		datas = append(datas, &data)

	}

	return datas
}

type UserResponseLogin struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

func (us *UserResponse) FormatterLoginResponse(user models.User) UserResponseLogin {
	return UserResponseLogin{
		Username:     user.Username,
		Email:        user.Email,
		Role:         ConvertRoleName(int(user.Roles)),
		AccessToken:  user.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: user.RefreshToken,
	}
}
