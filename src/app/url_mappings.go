package app

import "github.com/k-washi/example-golang-rest-api/src/controllers/health"

func mapUrls() {
	router.GET("/health", health.GetHealthStatusOK)
}
