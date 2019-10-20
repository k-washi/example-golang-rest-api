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
	getCollectInfoDaoFunction  func() (*collectinfo.GetCollectInfoResponse, errors.APIError)
)

func (m *collectInfoDaoMock) CreateInfoDao(input collectinfo.CollectInfo) errors.APIError {
	return postCollectInfoDaoFunction(input)
}

func (m *collectInfoDaoMock) GetInfoDao() (*collectinfo.GetCollectInfoResponse, errors.APIError) {
	return getCollectInfoDaoFunction()
}

func init() {
	collectinfo.CollectInfoDao = &collectInfoDaoMock{}
}

func TestCreateCollectInfoSuccess(t *testing.T) {
	postCollectInfoDaoFunction = func(input collectinfo.CollectInfo) errors.APIError {
		return nil
	}

	input := collectinfo.CreateCollectInfoRequest{
		Name:        "test",
		Description: "test desc",
		ID:          1,
		Data:        "test",
	}

	result, err := CollectInfoService.CreateCollectInfo(input)

	assert.Equal(t, result.Name, input.Name)
	assert.Equal(t, result.Message, "create info and store database")
	assert.Nil(t, err)
}

func TestGetCollectInfoSuccess(t *testing.T) {
	getCollectInfoDaoFunction = func() (*collectinfo.GetCollectInfoResponse, errors.APIError) {
		return &collectinfo.GetCollectInfoResponse{
			Datas: []collectinfo.CollectInfo{{ID: 1, Name: "test1"}, {ID: 123, Name: "test2"}},
		}, nil
	}

	result, err := CollectInfoService.GetCollectInfo()

	assert.Nil(t, err)
	assert.Equal(t, result.Name, "name")
	assert.Equal(t, result.Description, "test")
	assert.Equal(t, result.Datas[0].ID, 1)
	assert.Equal(t, result.Datas[0].Name, "test1")

}
