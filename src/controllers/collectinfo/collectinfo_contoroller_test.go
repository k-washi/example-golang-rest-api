package collectinfo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/k-washi/example-golang-rest-api/src/domain/collectinfo"
	"github.com/k-washi/example-golang-rest-api/src/services"
	"github.com/k-washi/example-golang-rest-api/src/utils/errors"
	"github.com/stretchr/testify/assert"
)

type postCollectInfoServiceMock struct {
}

var (
	postCollectInfoServiceFunction func(collectinfo.CreateCollectInfoRequest) (*collectinfo.CreateCollectInfoResponse, errors.APIError)
)

func (m *postCollectInfoServiceMock) CreateCollectInfo(input collectinfo.CreateCollectInfoRequest) (*collectinfo.CreateCollectInfoResponse, errors.APIError) {
	return postCollectInfoServiceFunction(input)
}

func init() {
	services.CollectInfoService = &postCollectInfoServiceMock{}
}

func TestCreateInfoContorollerSuccess(t *testing.T) {
	contextName := `"test"`
	serviceResponse := collectinfo.CreateCollectInfoResponse{
		Name:    contextName,
		Message: "create info and store database",
	}
	postCollectInfoServiceFunction = func(collectinfo.CreateCollectInfoRequest) (*collectinfo.CreateCollectInfoResponse, errors.APIError) {
		return &serviceResponse, nil
	}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	jsonRequest := `{"name":` + contextName + `, "description": "test", "data": {"id": 1, "name": "test1" }}`
	request, _ := http.NewRequest(http.MethodPost, "/example-golang-rest-api", bytes.NewBuffer([]byte(jsonRequest)))
	c.Request = request

	CreateInfo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	//Conver the JSON response to a map
	var collectInfoJSONResponse collectinfo.CreateCollectInfoResponse
	err := json.Unmarshal([]byte(response.Body.String()), &collectInfoJSONResponse)

	//Grab the balue & whether or not it exists
	nameStr := collectInfoJSONResponse.Name
	messageStr := collectInfoJSONResponse.Message

	assert.Nil(t, err)
	assert.Equal(t, serviceResponse.Name, nameStr)
	assert.Equal(t, serviceResponse.Message, messageStr)

}
