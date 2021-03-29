package controllers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var test = Login{
	Username: "username",
	Password: "password",
}

func Signin(c *gin.Context) {
	var user Login

	// Check JSON Valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, "ERROR INVALID JSON FORMAT")
		c.Abort()
		return
	}

	// Find in Data base
	if test.Username != user.Username || test.Password != user.Password {
		c.JSON(401, "UnAuthorization USER or PASS Invalid")
		c.Abort()
		return
	}

	// Gen token

}

func Signup(c *gin.Context) {

}

func CreateToken(id uint64) (string, error) {

	// Create Claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 60).Unix(),
		"id":  id,
	}

	// Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
