package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.POST("/signin")
	r.POST("/signup")

	r.Run() // listen and serve on 0.0.0.0:8080
}
