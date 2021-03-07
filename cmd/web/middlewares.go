package main

import (
	"github.com/gin-gonic/gin"
)

func secureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Frame-Options", "deny")
	}
}
