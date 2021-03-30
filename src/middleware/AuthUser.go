package middles

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	Token string `header:"Authorization"`
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := authHeader{}

		// Check Authorization Header && Extract Token
		if err := c.BindHeader(&header); err != nil {
			c.JSON(http.StatusForbidden, "UnAuthorization")
			c.Abort()
			return
		}

		// Splite Authorization
		tokenKey := strings.Split(header.Token, "Bearer ")

		// Check Bearer
		if len(tokenKey) < 2 {
			c.JSON(http.StatusUnauthorized,
				" Must provide Authorization header with format `Bearer {token}")
			c.Abort()
			return
		}

		// Vertify Token
		token, err := jwt.Parse(tokenKey[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		// If Parse Token Error that meain expire
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		// Cast Claims to MapClaims for using Map
		claims := token.Claims.(jwt.MapClaims)

		// Pass Userid to next
		c.Set("userId", claims["id"])

		c.Next()
	}
}
