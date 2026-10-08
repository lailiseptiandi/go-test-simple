package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lailiseptiandi/go-test-simple/internal/dtos/responses"
	"github.com/lailiseptiandi/go-test-simple/internal/models"
)

var jwtSercret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Name     string `json:"name"`
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(user *models.User) (string, error) {
	// JWT Expired
	jwtTokenExpiresAt, _ := time.ParseDuration(os.Getenv("JWT_EXPIRATION") + "s")

	claims := Claims{
		UserID:   uint64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Role:     responses.ConvertRoleName(int(user.Roles)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtTokenExpiresAt)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSercret)
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return tokenString, nil
}

func GenerateRefreshToken(user *models.User) (string, error) {
	// JWT Expired
	jwtRefreshTokenExpiresAt, _ := time.ParseDuration(os.Getenv("JWT_EXPIRATION_REFRESH") + "s")

	claims := Claims{
		UserID:   uint64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Role:     responses.ConvertRoleName(int(user.Roles)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtRefreshTokenExpiresAt)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshTokenString, err := refreshToken.SignedString(jwtSercret)
	if err != nil {
		return "", errors.New("failed to generate refresh token")
	}

	return refreshTokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return jwtSercret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return claims, nil

}
