package authmiddle

import (
	"net/http"
	"strings"

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
			c.JSON(http.StatusUnauthorized, "UnAuthorization")
			c.Abort()
			return
		}

		// Splite Authorization
		tokenId := strings.Split(header.Token, "Bearer ")

		// Check Bearer
		if len(tokenId) < 2 {
			c.JSON(http.StatusUnauthorized,
				" Must provide Authorization header with format `Bearer {token}")
			c.Abort()
			return
		}

		c.Next()
	}
}
