package routers

import (
	"Auth/src/controllers"
	"Auth/src/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RouterAuth(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.Use(authmiddle.AuthUser())
	{
		auth.GET("/check", func(c *gin.Context) {
			fmt.Println(c.MustGet("userId").(string))
		})
	}
	auth.POST("/signin", controllers.Signin)

}
