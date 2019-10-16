package health

import (
	"net/http"

	"github.com/k-washi/example-golang-rest-api/src/services"

	"github.com/gin-gonic/gin"
)

//GetHealthStatusOK status 200 is responsed.
func GetHealthStatusOK(c *gin.Context) {

	/* //No json request
	var request repositories.CreateHealthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}
	*/

	result, err := services.HealthService.GetHealth()
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}
	c.JSON(http.StatusOK, result)

}
