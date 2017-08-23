package app

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/syedomair/weather/testdata"
)

func TestPostValidAddressAction(t *testing.T) {
	e := httpexpect.New(t, server.URL)

	address := map[string]interface{}{
		"address": testdata.ValidAddress1,
	}
	p := e.POST("/v1/weather").WithJSON(address).Expect()
	o := p.Status(http.StatusOK).JSON().Object()

	o.ValueEqual("result", "success")
}

func TestPostInvalidAddressAction(t *testing.T) {
	e := httpexpect.New(t, server.URL)

	address := map[string]interface{}{
		"address": testdata.InValidAddress1,
	}
	p := e.POST("/v1/weather").WithJSON(address).Expect()
	o := p.Status(http.StatusBadRequest).JSON().Object()

	o.ValueEqual("result", "error")
}
