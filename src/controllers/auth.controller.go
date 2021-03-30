package controllers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

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

type ResultQuery struct {
	ID       uint64
	Email    string
	Username string
	Password string
	Fname    string
	Lname    string
	RoleID   uint
}

var test = Login{
	Username: "username",
	Password: "password",
}

func Signin(c *gin.Context) {
	var ctx = c.MustGet("dbConnection").(*gorm.DB)
	var user Login
	var retriveUser models.User
	// Check JSON Valid
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusAccepted, "ERROR INVALID JSON FORMAT")
		c.Abort()
		return
	}

	// Find in Data base
	result := ctx.Debug().Preload("Roles").Where("username = ?", user.Username).First(&retriveUser)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found ...."})
		c.Abort()
		return
	}

	// Compare Password
	if err := decryptPassword(user.Password, retriveUser.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Password incorrect"})
		c.Abort()
		return
	}

	// Gen token
	token, err := CreateToken(retriveUser.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Add Role
	role := []string{}

	for _, val := range retriveUser.Roles {
		role = append(role, "ROLE_"+strings.ToUpper(val.Role))
	}

	c.JSON(http.StatusOK, gin.H{
		"firstname": retriveUser.Fname,
		"lastname":  retriveUser.Lname,
		"token":     token,
		"role":      strings.Join(role, ","),
	})

}

func Signup(c *gin.Context) {
	var ctx = c.MustGet("dbConnection").(*gorm.DB)
	var user models.User
	var params = c.MustGet("params").(SignupType) // Get Params from MiddleWare

	user.Email = params.Email
	user.Username = params.Username
	user.Password = string(genBcrypt(params.Password)) // bcrypt password hashing
	user.Fname = params.Fname
	user.Lname = params.Lname
	user.Roles = []models.Role{{ID: 1}, {ID: 2}}

	if result := ctx.Create(&user); result.Error != nil {
		c.JSON(http.StatusForbidden, result.Error)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User was registered successfully!"})

}

func CreateToken(id uint) (string, error) {

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

// Compare Password
func decryptPassword(rawPass string, hashedPassword string) error {

	// Succuss return nil , if failure return error
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPass))

	return err
}
