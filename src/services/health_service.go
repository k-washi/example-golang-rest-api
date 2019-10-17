package services

import (
	"net/http"

	"github.com/k-washi/example-golang-rest-api/src/domain/health"
	"github.com/k-washi/example-golang-rest-api/src/utils/errors"
)

type healthServiceInterface interface {
	GetHealth() (*health.CreateHealthResponse, errors.APIError)
}

var (
	//HealthService health check service
	HealthService healthServiceInterface
)

type healthService struct{}

func init() {
	HealthService = &healthService{}
}

func (s *healthService) GetHealth() (*health.CreateHealthResponse, errors.APIError) {
	result := health.CreateHealthResponse{
		Status:      http.StatusOK,
		Description: "Health check OK",
	}
	return &result, nil
}
