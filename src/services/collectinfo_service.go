package services

import (
	"strings"

	"github.com/k-washi/example-golang-rest-api/src/domain/collectinfo"
	"github.com/k-washi/example-golang-rest-api/src/utils/errors"
)

type collectInfoServiceInterface interface {
	CreateCollectInfo(collectinfo.CreateCollectInfoRequest) (*collectinfo.CreateCollectInfoResponse, errors.APIError)
}

var (
	//CollectInfoService for access domain
	CollectInfoService collectInfoServiceInterface
)

type collectInfoService struct{}

func init() {
	CollectInfoService = &collectInfoService{}
}

//StoreCollectInfo store info and return result
func (s *collectInfoService) CreateCollectInfo(input collectinfo.CreateCollectInfoRequest) (*collectinfo.CreateCollectInfoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid collection info name")
	}

	err := collectinfo.CollectInfoDao.CreateInfoDao(input.Data)
	if err != nil {
		return nil, errors.NewAPIError(err.GetStatus(), err.GetMessage())
	}

	result := collectinfo.CreateCollectInfoResponse{
		Name:    input.Name,
		Message: "create info and store database",
	}
	return &result, nil
}
