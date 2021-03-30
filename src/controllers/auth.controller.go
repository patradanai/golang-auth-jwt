package controllers

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"Auth/src/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupType struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
}

var test = Login{
	Username: "username",
	Password: "password",
}

func Signin(c *gin.Context) {
	var user Login

	// Check JSON Valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusAccepted, "ERROR INVALID JSON FORMAT")
		c.Abort()
		return
	}

	// Find in Data base
	if test.Username != user.Username || test.Password != user.Password {
		c.JSON(http.StatusUnauthorized, "UnAuthorization USER or PASS Invalid")
		c.Abort()
		return
	}

	// Gen token
	token, err := CreateToken(1)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func Signup(c *gin.Context) {
	var user models.User
	var params = c.MustGet("params").(SignupType)

	user.Email = params.Email
	user.Username = params.Username
	user.Password = string(genBcrypt(params.Password)) // bcrypt password hashing
	user.Fname = params.Fname
	user.Lname = params.Lname
	user.RoleID = 2 // Customer

	result := models.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusForbidden, result.Error)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User was registered successfully!"})

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

// Hashing Password
func genBcrypt(rawPass string) []byte {
	var password = []byte(rawPass)

	// Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hashedPassword
}
