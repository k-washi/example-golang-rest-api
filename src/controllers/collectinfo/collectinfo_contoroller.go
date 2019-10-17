package collectinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k-washi/example-golang-rest-api/src/domain/collectinfo"
	"github.com/k-washi/example-golang-rest-api/src/services"
	"github.com/k-washi/example-golang-rest-api/src/utils/errors"
)

//CreateInfo collect {id, name} data
//req: {name: "name", description: "test", data: {id: 1, name: "test1" }}, res: {name: "name", message: "ok"}
func CreateInfo(c *gin.Context) {
	var request collectinfo.CreateCollectInfoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	result, err := services.CollectInfoService.CreateCollectInfo(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)

}
