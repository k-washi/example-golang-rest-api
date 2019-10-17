package collectinfo

import (
	"reflect"

	"github.com/k-washi/example-golang-rest-api/src/utils/errors"

	"log"
)

type collectInfoInterface interface {
	CreateInfoDao(CollectInfo) errors.APIError
	GetInfoDao() (*GetCollectInfoResponse, errors.APIError)
}

var (
	collectInfos = map[int]CollectInfo{
		123: {ID: 123, Name: "test-name"},
	}

	//CollectInfoDao collect info interface
	CollectInfoDao collectInfoInterface
)

type collectInfoDao struct{}

func init() {
	CollectInfoDao = &collectInfoDao{}
}

func (d *collectInfoDao) CreateInfoDao(input CollectInfo) errors.APIError {
	/*
	 */
	//evaluate type int
	if reflect.TypeOf(input.ID).Kind() != reflect.Int {
		return errors.NewBadData("input id type mismatch")
	}

	if _, ok := collectInfos[input.ID]; ok {
		//データがすでに存在する
		log.Printf("input data exsist")
	}
	collectInfos[input.ID] = input

	return nil

}

func (d *collectInfoDao) GetInfoDao() (*GetCollectInfoResponse, errors.APIError) {
	var collectInfoList GetCollectInfoResponse
	for _, v := range collectInfos {
		collectInfoList.Datas = append(collectInfoList.Datas, v)
	}

	return &collectInfoList, nil

}
