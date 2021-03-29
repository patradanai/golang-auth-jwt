package authmiddle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type authHeader struct {
	token string `header:"Authorization"`
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := authHeader{}

		// Check Authorization Header
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(401, "UnAuthorization")
			c.Abort()
			return
		}

		// Splite Authorization
		tokenId := strings.Split(header.token, "Bearer ")
		fmt.Printf("%#v\n", header)
		if len(tokenId) < 2 {
			c.JSON(401,
				" Must provide Authorization header with format `Bearer {token}")
			c.Abort()
			return
		}

		c.Next()
	}
}
