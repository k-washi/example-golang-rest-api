package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"

	"github.com/k-washi/example-golang-rest-api/src/domain/health"
	"github.com/k-washi/example-golang-rest-api/src/services"
	"github.com/k-washi/example-golang-rest-api/src/utils/errors"
)

/*
#依存関係のあるhealthService.GetHealthのモックを作成
*/
type healthServiceMock struct {
}

var (
	getHealthFunction func() (*health.CreateHealthResponse, errors.ApiError)
)

func (m *healthServiceMock) GetHealth() (*health.CreateHealthResponse, errors.ApiError) {
	return getHealthFunction()
}

func init() {
	services.HealthService = &healthServiceMock{}
}

//TestGetHealthStatusOK test of status 200 with service mock
func TestGetHealthStatusOK(t *testing.T) {
	exsistHealthResponse := health.CreateHealthResponse{
		Status:      http.StatusOK,
		Description: "Health check OK",
	}

	getHealthFunction = func() (*health.CreateHealthResponse, errors.ApiError) {
		return &exsistHealthResponse, nil
	}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	request, _ := http.NewRequest(http.MethodGet, "/health", nil)
	c.Request = request

	GetHealthStatusOK(c)

	assert.EqualValues(t, http.StatusOK, response.Code)

	//Conver the JSON response to a map
	var healthJSONResponse health.CreateHealthResponse
	err := json.Unmarshal([]byte(response.Body.String()), &healthJSONResponse)

	//Grab the balue & whether or not it exists
	statusVal := healthJSONResponse.Status
	descriptionVal := healthJSONResponse.Description

	assert.Nil(t, err)
	assert.Equal(t, exsistHealthResponse.Status, statusVal)
	assert.Equal(t, exsistHealthResponse.Description, descriptionVal)

}
