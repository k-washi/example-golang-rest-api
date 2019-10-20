package app

import (
	"github.com/gin-gonic/gin"
	"github.com/k-washi/example-golang-rest-api/src/middleware"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

//StartApp call by main for starting app.
func StartApp() {
	router.Use(middleware.OptionsMethodResponse())
	router.Use(middleware.HeaderSet())
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
