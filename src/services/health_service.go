package services

import (
	"net/http"

	"github.com/k-washi/example-golang-rest-api/src/domain/repositories"
	"github.com/k-washi/example-golang-rest-api/src/utils/errors"
)

type healthServiceInterface interface {
	CreateRepo() (*repositories.CreateHealthResponse, errors.ApiError)
}

var (
	HealthService healthServiceInterface
)

type healthService struct{}

func init() {
	HealthService = &healthService{}
}

func (s *healthService) CreateRepo() (*repositories.CreateHealthResponse, errors.ApiError) {
	result := repositories.CreateHealthResponse{
		Status:      http.StatusOK,
		Description: "Health check OK",
	}
	return &result, nil
}
