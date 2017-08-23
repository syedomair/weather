package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-kit/kit/log"
	"github.com/syedomair/weather/app/common"
	"github.com/syedomair/weather/models"
)

type DataResponse struct {
	Address     string `json:"address"`
	Temperature string `json:"temperature"`
}

func PostAction(c *gin.Context, config common.Config, logger log.Logger) {

	logger.Log("action", "PostAction", "event", "method start")
	start := time.Now()

	address, err := GetAddressFromRequest(c.Request.Body)
	if err != nil {
		str := err.Error()
		c.JSON(http.StatusBadRequest, common.ErrorResponse(str))
		return
	}

	elapsed := time.Since(start)
	logger.Log("action", "PostAction", "GetAddressFromRequest", address, "time_spent", elapsed)
	start = time.Now()

	lng, lat, f_address, err := GetCoordinateFromAddress(config, address)
	if err != nil {
		str := err.Error()
		c.JSON(http.StatusBadRequest, common.ErrorResponse(str))
		return
	}
	elapsed = time.Since(start)
	logger.Log("action", "PostAction", "GetCoordinateFromAddress", f_address, "time_spent", elapsed)

	temperature_chan := make(chan string)

	go func() {
		start = time.Now()
		logger.Log("action", "PostAction", "GetTemperatureFromCoordinate", "", "time_start", start)
		temperature, err := GetTemperatureFromCoordinate(config, lng, lat)
		if err != nil {
			str := err.Error()
			c.JSON(http.StatusBadRequest, common.ErrorResponse(str))
			return
		}
		elapsed = time.Since(start)
		logger.Log("action", "PostAction", "GetTemperatureFromCoordinate", temperature, "time_spent", elapsed)
		temperature_chan <- temperature
	}()

	go func() {
		start = time.Now()
		logger.Log("action", "PostAction", "PostToWeatherLogRecord", "", "time_start", start)
		dbInterface, _ := c.Get("DB")
		db := dbInterface.(models.Datastore)

		err = PostToWeatherLogRecord(db, c.ClientIP(), f_address)
		if err != nil {
			str := err.Error()
			c.JSON(http.StatusBadRequest, common.ErrorResponse(str))
			return
		}

		elapsed = time.Since(start)
		logger.Log("action", "PostAction", "PostToWeatherLogRecord", "", "time_spent", elapsed)

	}()

	var dataResponse DataResponse
	dataResponse.Address = f_address
	dataResponse.Temperature = <-temperature_chan

	c.JSON(http.StatusOK, common.SuccessResponse(dataResponse))
}
