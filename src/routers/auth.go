package routers

import (
	"Auth/src/controllers"

	"github.com/gin-gonic/gin"
)

func RouterAuth(router *gin.Engine) {
	auth := router.Group("/auth")

	auth.POST("/signin", controllers.Signin)

}
