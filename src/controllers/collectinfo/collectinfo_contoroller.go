package collectinfo

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k-washi/example-golang-rest-api/src/domain/collectinfo"
	"github.com/k-washi/example-golang-rest-api/src/services"
	"github.com/k-washi/example-golang-rest-api/src/utils/errors"
)

//CreateInfo collect {id, name} data
//POST: req: {"name": "name", "description": "test", "id": 1, "data": "test1" }, res: {"name": "name", "message": "c.."}
func CreateInfo(c *gin.Context) {
	var request collectinfo.CreateCollectInfoRequest

	if err := c.ShouldBind(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}
	name := c.PostForm("name")
	data := c.PostForm("data")
	log.Printf("%s, %s", name, data)
	log.Printf("%s, %s", request.Name, request.Data)

	result, err := services.CollectInfoService.CreateCollectInfo(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)

}

//GetInfo get [{id, name}] datas
//req: ?name=name, res: {"name": "name", "description": "test", "datas": [{post-data-1}, {post-data-2}] }
func GetInfo(c *gin.Context) {
	//var request collectinfo.GetCollectInfoRequest
	/* /~/:nameで設定された値を取得
	param := c.Param("name")
	*/
	param := c.Query("name")
	log.Println(param)
	/*
		if err := c.ShouldBindJSON(&request); err != nil {
			apiErr := errors.NewBadRequestError("invalid json body")
			c.JSON(apiErr.GetStatus(), apiErr)
			return
		}
	*/

	//TODO
	result, err := services.CollectInfoService.GetCollectInfo()

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, result)

}
