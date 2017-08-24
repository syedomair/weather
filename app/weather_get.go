package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syedomair/weather/common"
	"github.com/syedomair/weather/models"
)

func WeatherGetAction(c *gin.Context) {

	dbInterface, _ := c.Get("DB")
	db := dbInterface.(models.Datastore)

	wls, err := GetAllWeatherLogRecords(db)
	if err != nil {
		str := err.Error()
		c.JSON(http.StatusBadRequest, common.ErrorResponse(str))
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(wls))
}
