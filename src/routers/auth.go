package routers

import (
	middles "Auth/src/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RouterAuth(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.Use(middles.AuthUser())
	{
		auth.GET("/check", func(c *gin.Context) {
			fmt.Println(c.MustGet("userId").(string))
		})
	}
}
