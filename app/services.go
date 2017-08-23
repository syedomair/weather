package app

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/syedomair/weather/app/common"
	"github.com/syedomair/weather/models"
)

func PostToWeatherLogRecord(db models.Datastore, client_ip string, address string) error {

	var weatherLog models.WeatherLog
	weatherLog.IpAddress = client_ip
	weatherLog.AddressSearched = address

	err := db.PostWeatherLog(weatherLog)
	if err != nil {
		err := errors.New(Error_code_10007)
		return err
	}

	return nil
}
func GetTemperatureFromCoordinate(config common.Config, longitude string, latitude string) (string, error) {

	completeURL := config.DarkskyURL + config.DarkskyKey + "/" + latitude + "," + longitude

	mapBodyInterface, err := getDataFromThirdPartyServer(completeURL)
	if err != nil {
		err := errors.New(Error_code_10006)
		return "", err
	}
	currently := make(map[string]interface{})
	if _, ok := mapBodyInterface["code"]; ok {
		err := errors.New(Error_code_10006)
		return "", err
	}
	currently = mapBodyInterface["currently"].(map[string]interface{})
	temperature := strconv.FormatFloat(currently["temperature"].(float64), 'f', -1, 64)
	return temperature, nil
}

func GetCoordinateFromAddress(config common.Config, address string) (string, string, string, error) {
	var urlLocal *url.URL
	urlLocal, err := url.Parse(config.MapsGoogleURL)
	if err != nil {
		err := errors.New(Error_code_10003)
		return "", "", "", err
	}
	parameters := url.Values{}
	parameters.Add("address", address)
	urlLocal.RawQuery = parameters.Encode()

	completeURL := urlLocal.String()

	mapBodyInterface, err := getDataFromThirdPartyServer(completeURL)
	if err != nil {
		err := errors.New(Error_code_10004)
		return "", "", "", err
	}

	results := mapBodyInterface["results"]
	status := mapBodyInterface["status"]
	longitute, latitude, f_address := "", "", ""
	if status == "OK" {
		resultsArray := results.([]interface{})

		resultElement := make(map[string]interface{})
		resultElement = resultsArray[0].(map[string]interface{})

		geometry := make(map[string]interface{})
		geometry = resultElement["geometry"].(map[string]interface{})

		location := make(map[string]interface{})
		location = geometry["location"].(map[string]interface{})

		latitude = strconv.FormatFloat(location["lat"].(float64), 'f', -1, 64)
		longitute = strconv.FormatFloat(location["lng"].(float64), 'f', -1, 64)
		f_address = resultElement["formatted_address"].(string)
	} else {
		err := errors.New(Error_code_10005)
		return "", "", "", err
	}
	return longitute, latitude, f_address, nil
}

func getDataFromThirdPartyServer(completeURL string) (map[string]interface{}, error) {
	resp, err := http.Get(completeURL)
	if err != nil {
		err := errors.New(Error_code_10002)
		return nil, err
	}
	defer resp.Body.Close()
	mapBody, err := ioutil.ReadAll(resp.Body)

	mapBodyInterface := make(map[string]interface{})
	json.Unmarshal(mapBody, &mapBodyInterface)
	return mapBodyInterface, nil
}

type InputData struct {
	Address string `json:"address"`
}

func GetAddressFromRequest(reqBody io.ReadCloser) (string, error) {
	body, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()
	if err != nil {
		err := errors.New(Error_code_10001)
		return "", err
	}
	var bodyInterface map[string]interface{}
	json.Unmarshal(body, &bodyInterface)

	var inputData InputData
	err = json.Unmarshal(body, &inputData)
	if err != nil {
		err := errors.New(Error_code_10001)
		return "", err
	}
	if inputData.Address == "" {
		err := errors.New(Error_code_10001)
		return "", err
	}
	address := inputData.Address
	return address, nil
}
