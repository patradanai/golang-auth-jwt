package middles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Auth/src/controllers"
	"Auth/src/models"
)

func VertifySignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.MustGet("dbConnection").(*gorm.DB)
		// Slice have child User
		var user []models.User
		// paramsUser
		var paramsUser controllers.SignupType

		// Extract ParamsUser
		if err := c.ShouldBindJSON(&paramsUser); err != nil {
			c.JSON(http.StatusForbidden, "Format Invalid")
			c.Abort()
			return
		}

		// Find in DB
		result := ctx.Where("username = ?", paramsUser.Username).First(&user)

		// Check If existing User
		if result.RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed! Username is already in use!"})
			c.Abort()
			return
		}

		c.Set("params", paramsUser)

		c.Next()
	}
}
