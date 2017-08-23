package app

import (
	"testing"

	"github.com/syedomair/weather/testdata"
)

func TestGetTemperatureFromCoordinate(t *testing.T) {

	temp, _ := GetTemperatureFromCoordinate(config, testdata.ValidLatitude, testdata.ValidLogitute)

	if temp == "" {
		t.Error("Expected a tempreture")
	}

	temp, _ = GetTemperatureFromCoordinate(config, testdata.InValidLatitude, testdata.InValidLogitute)
	if temp != "" {
		t.Error("Expected an string")
	}
}

func TestGetCoordinateFromAddress(t *testing.T) {

	_, _, f_address, _ := GetCoordinateFromAddress(config, testdata.ValidAddress1)
	if f_address == "" {
		t.Error("Expected an address")
	}

	_, _, f_address, _ = GetCoordinateFromAddress(config, testdata.InValidAddress1)
	if f_address != "" {
		t.Error("Expected an empty address")
	}
}
