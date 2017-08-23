package models

import (
	"fmt"
	"time"
)

type WeatherLog struct {
	Id              string
	IpAddress       string
	AddressSearched string
	CreatedAt       time.Time
}

func (db *DB) PostWeatherLog(w WeatherLog) error {

	sql := fmt.Sprintf("INSERT INTO weather_log (ip_address, address_searched, created_at) VALUES('%s', '%s',CURRENT_TIMESTAMP )", w.IpAddress, w.AddressSearched)

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil

}
