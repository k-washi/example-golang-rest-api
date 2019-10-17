package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/k-washi/example-golang-rest-api/src/domain/collectinfo"
	"github.com/k-washi/example-golang-rest-api/src/utils/errors"
)

type collectInfoDaoMock struct {
}

var (
	postCollectInfoDaoFunction func(collectinfo.CollectInfo) errors.APIError
)

func (m *collectInfoDaoMock) CreateInfoDao(input collectinfo.CollectInfo) errors.APIError {
	return postCollectInfoDaoFunction(input)
}

func init() {
	collectinfo.CollectInfoDao = &collectInfoDaoMock{}
}

func TestCreateCollectInfoSuccess(t *testing.T) {
	postCollectInfoDaoFunction = func(input collectinfo.CollectInfo) errors.APIError {
		return nil
	}

	/*
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)

		jsonRequest := `{"name": "name", "description": "test", "data": {"id": 1, "name": "test1" }}`
		request, _ := http.NewRequest(http.MethodPost, "/example-golang-rest-api", bytes.NewBuffer([]byte(jsonRequest)))
		c.Request = request
	*/

	collectData := collectinfo.CollectInfo{
		ID:   1,
		Name: "test",
	}
	input := collectinfo.CreateCollectInfoRequest{
		Name:        "test",
		Description: "test desc",
		Data:        collectData,
	}

	result, err := CollectInfoService.CreateCollectInfo(input)

	assert.Equal(t, result.Name, input.Name)
	assert.Equal(t, result.Message, "create info and store database")
	assert.Nil(t, err)
}
