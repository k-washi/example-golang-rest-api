package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//OptionsMethodResponse CROS options response
func OptionsMethodResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.Header().Set("Content-Type", "application/json")
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			c.Abort()
		}
		return

	}
}
