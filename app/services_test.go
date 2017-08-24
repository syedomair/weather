package app

import (
	"bytes"
	"io"
	"testing"

	"github.com/syedomair/weather/mockdata"
)

func TestGetAllWeatherLogRecords(t *testing.T) {
	wls, _ := GetAllWeatherLogRecords(db)
	if wls[0].AddressSearched == "" {
		t.Error("Expected an address")
	}
}
func TestGetTemperatureFromCoordinate(t *testing.T) {

	temp, _ := GetTemperatureFromCoordinate(config, mockdata.ValidLatitude, mockdata.ValidLogitute)

	if temp == "" {
		t.Error("Expected a tempreture")
	}

	temp, _ = GetTemperatureFromCoordinate(config, mockdata.InValidLatitude, mockdata.InValidLogitute)
	if temp != "" {
		t.Error("Expected an string")
	}
}

func TestGetCoordinateFromAddress(t *testing.T) {

	_, _, f_address, _ := GetCoordinateFromAddress(config, mockdata.ValidAddress1)
	if f_address == "" {
		t.Error("Expected an address")
	}

	_, _, f_address, _ = GetCoordinateFromAddress(config, mockdata.InValidAddress1)
	if f_address != "" {
		t.Error("Expected an empty address")
	}
}

func TestPostToWeatherLogRecord(t *testing.T) {

	id, err := PostToWeatherLogRecord(db, mockdata.IpAddress, mockdata.TestAddress)

	if err != nil {
		t.Error("Error while posting record to the database")
	}

	wl, err := db.GetAWeatherLog(id)
	if err != nil {
		t.Error("Error while fetching record from the database")
	}
	if wl.AddressSearched == "" {
		t.Error("Expected an address")
	}
}

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb *ClosingBuffer) Close() (err error) {
	return
}

func TestGetAddressFromRequest(t *testing.T) {

	cb := &ClosingBuffer{bytes.NewBufferString(mockdata.ValidJSONAddress)}
	var reqBody io.ReadCloser
	reqBody = cb
	defer reqBody.Close()
	address, _ := GetAddressFromRequest(reqBody)
	if address == "" {
		t.Error("Expected an address")
	}
}
