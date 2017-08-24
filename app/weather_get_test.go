package app

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestWeatherGetAction(t *testing.T) {
	e := httpexpect.New(t, server.URL)

	p := e.GET("/v1/weather-log").Expect()
	o := p.Status(http.StatusOK).JSON().Object()

	o.ValueEqual("result", "success")
}
