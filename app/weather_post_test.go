package app

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/syedomair/weather/mockdata"
)

func TestWeatherPostValidAddressAction(t *testing.T) {
	e := httpexpect.New(t, server.URL)

	address := map[string]interface{}{
		"address": mockdata.ValidAddress1,
	}
	p := e.POST("/v1/weather").WithJSON(address).Expect()
	o := p.Status(http.StatusOK).JSON().Object()

	o.ValueEqual("result", "success")
}

func TestWeatherPostInvalidAddressAction(t *testing.T) {
	e := httpexpect.New(t, server.URL)

	address := map[string]interface{}{
		"address": mockdata.InValidAddress1,
	}
	p := e.POST("/v1/weather").WithJSON(address).Expect()
	o := p.Status(http.StatusBadRequest).JSON().Object()

	o.ValueEqual("result", "error")
}
