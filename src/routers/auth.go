package routers

import (
	"Auth/src/controllers"
	middles "Auth/src/middleware"

	"github.com/gin-gonic/gin"
)

func RouterAuth(router *gin.Engine) {
	auth := router.Group("/auth")
	// auth.Use(middles.AuthUser())
	// {
	// 	auth.GET("/check", func(c *gin.Context) {
	// 		fmt.Println(c.MustGet("userId").(string))
	// 	})
	// }
	auth.POST("/signin", controllers.Signin)
	auth.POST("/signup", middles.VertifySignUp(), controllers.Signup)

}
