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

func (db *DB) PostWeatherLog(w WeatherLog) (string, error) {

	sql := fmt.Sprintf("INSERT INTO weather_log (ip_address, address_searched, created_at) VALUES('%s', '%s',CURRENT_TIMESTAMP ) returning id;", w.IpAddress, w.AddressSearched)

	var id string
	err := db.QueryRow(sql).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil

}

func (db *DB) GetAllWeatherLog() ([]*WeatherLog, error) {

	rows, err := db.Query("SELECT * FROM weather_log")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	wls := make([]*WeatherLog, 0)
	for rows.Next() {
		wl := new(WeatherLog)
		err := rows.Scan(&wl.Id, &wl.IpAddress, &wl.AddressSearched, &wl.CreatedAt)
		if err != nil {
			return nil, err
		}
		wls = append(wls, wl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return wls, nil

}
func (db *DB) GetAWeatherLog(id string) (*WeatherLog, error) {
	wl := new(WeatherLog)
	stmt, err := db.Prepare("select id, ip_address, address_searched, created_at  from weather_log where id = $1 ")
	if err != nil {
		return wl, err
	}
	err = stmt.QueryRow(id).Scan(&wl.Id, &wl.IpAddress, &wl.AddressSearched, &wl.CreatedAt)
	if err != nil {
		return wl, err
	}
	return wl, nil
}
