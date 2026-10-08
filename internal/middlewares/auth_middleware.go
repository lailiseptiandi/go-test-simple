package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/jwt"
	"github.com/lailiseptiandi/go-test-simple/pkg/utils/response"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Unauthorized: Token is required")
			return
		}

		tokenString := ""
		authToken := strings.Split(tokenHeader, " ")
		if len(authToken) == 2 {
			tokenString = authToken[1]
		}

		parseToken, err := jwt.ParseToken(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Unauthorized: Invalid token")
			return
		}

		c.Set("credential_user", parseToken)
		c.Next()
	}
}
