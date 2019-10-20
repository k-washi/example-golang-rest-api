package middleware

import (
	"github.com/gin-gonic/gin"
)

//HeaderSet return header setting middleware function
func HeaderSet() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		 */
		//CORS header
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}
