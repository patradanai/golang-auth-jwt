package main

import (
	"Auth/src/models"
	"Auth/src/routers"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type authHeader struct {
	token string `header:"Authorization"`
}

// Create some middleware which swaps out the existing request context
// with new context.Context value containing the connection pool.
func injectDBGorm(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConnection", db)

		c.Next()
	}
}

func main() {

	// Please define your username and password for MySQL.
	DB, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}

	// Init Gorm database
	init := models.Handler{}
	init.Db = DB
	init.OpenDb()

	r := gin.Default()

	r.Use(injectDBGorm(DB))

	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Adding Cors
	r.Use(CORSMiddleware())

	// Initial Router
	routers.RouterAuth(r)
	routers.RouterSinging(r)

	// handle Timeout
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
