package models

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Datastore interface {
	PostWeatherLog(WeatherLog) (string, error)
	GetAWeatherLog(string) (*WeatherLog, error)
	GetAllWeatherLog() ([]*WeatherLog, error)
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	//db, err := sql.Open("postgres", dataSourceName)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
