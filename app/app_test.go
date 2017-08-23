package app

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/syedomair/weather/app/common"

	"net/http/httptest"

	"os"

	"path/filepath"
)

var server *httptest.Server
var config common.Config

func TestMain(m *testing.M) {
	server, config = testServer()
	defer server.Close()
	retCode := m.Run()
	os.Exit(retCode)
}

func testServer() (*httptest.Server, common.Config) {
	d, _ := os.Getwd()

	app := CreateGinApplication(gin.TestMode, filepath.Join(filepath.Dir(d), "config", "config_test.yml"))

	return httptest.NewServer(app.Engine), app.config
}
