package app

import (
	"github.com/k-washi/example-golang-rest-api/src/controllers/collectinfo"
	"github.com/k-washi/example-golang-rest-api/src/controllers/health"
)

func mapUrls() {
	router.GET("/health", health.GetHealthStatusOK)
	router.POST("/example-golang-rest-api", collectinfo.CreateInfo)
	router.GET("/example-golang-rest-api", collectinfo.GetInfo)

}
