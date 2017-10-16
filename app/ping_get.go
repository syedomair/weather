package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syedomair/weather/common"
)

func PingGetAction(c *gin.Context) {

	c.JSON(http.StatusOK, common.SuccessResponse("Pong"))
}
