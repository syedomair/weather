package app

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestPingGetAction(t *testing.T) {
	e := httpexpect.New(t, server.URL)

	p := e.GET("/v1/ping").Expect()
	o := p.Status(http.StatusOK).JSON().Object()

	o.ValueEqual("result", "success")
}
