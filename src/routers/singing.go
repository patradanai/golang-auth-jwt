package routers

import (
	"Auth/src/controllers"
	middles "Auth/src/middleware"

	"github.com/gin-gonic/gin"
)

func RouterSinging(router *gin.Engine) {
	router.POST("/signin", controllers.Signin)
	router.POST("/signup", middles.VertifySignUp(), controllers.Signup)
}
