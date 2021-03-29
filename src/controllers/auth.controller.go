package controllers

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthLogin(c *gin.Context) {
	var user Login

	// Check JSON Valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, "ERROR INVALID JSON FORMAT")
		c.Abort()
		return
	}

}
