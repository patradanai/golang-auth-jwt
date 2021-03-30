package middles

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Auth/src/models"
)

func VertifySignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Slice have child User
		var user []models.User
		// Find in DB
		result := models.DB.Where("username = ?", "patradanai").Find(&user)

		// Check If existing User
		if result.RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed! Username is already in use!"})
			c.Abort()
			return
		}

		c.Next()
	}
}
