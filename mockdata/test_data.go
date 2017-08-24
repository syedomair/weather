package mockdata

import "github.com/syedomair/weather/models"

func init() {
	AddWeatherData(models.WeatherLog{
		Id:              WeatherLogId,
		IpAddress:       WeatherLogIpAddress,
		AddressSearched: WeatherLogAddress})
}
