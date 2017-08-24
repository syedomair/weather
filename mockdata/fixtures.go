package mockdata

import (
	"github.com/syedomair/weather/models"
)

type record struct {
	Collection string
	Object     interface{}
}

type key struct {
	Collection string
	Key        string
}

var mdata = make(map[key]record)

func AddWeatherData(weatherLog models.WeatherLog) {
	mdata[key{Collection: "Id", Key: weatherLog.Id}] = record{Collection: "WeatherLog", Object: weatherLog}
}

func GetWeatherData(weatherLogId string) models.WeatherLog {
	return mdata[key{Collection: "Id", Key: weatherLogId}].Object.(models.WeatherLog)
}
