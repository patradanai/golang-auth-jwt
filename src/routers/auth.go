package routers

import (
	"Auth/src/controllers"
	"Auth/src/middleware"
	"github.com/gin-gonic/gin"
)

func RouterAuth(router *gin.Engine) {
	auth := router.Group("/auth")

	auth.GET("/check", authmiddle.AuthUser())
	auth.POST("/signin", controllers.Signin)

}
