package main

import "github.com/gin-gonic/gin"

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Authorization Header

		// Splite Authorization

		c.Next()
	}
}
