package app

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/syedomair/weather/common"
	"github.com/syedomair/weather/models"

	"net/http/httptest"

	"os"
	"path/filepath"

	"github.com/go-kit/kit/log"
)

var server *httptest.Server
var config common.Config
var logger log.Logger
var db models.Datastore

func TestMain(m *testing.M) {
	server, config, logger, db = testServer()
	defer server.Close()
	retCode := m.Run()
	os.Exit(retCode)
}

func testServer() (*httptest.Server, common.Config, log.Logger, models.Datastore) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "time", log.DefaultTimestamp)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	d, _ := os.Getwd()

	app := CreateGinApplication(gin.TestMode, filepath.Join(filepath.Dir(d), "config", "config_test.yml"), logger)

	logger.Log("transport", "HTTP", "addr", app.Config.HttpAddress)

	return httptest.NewServer(app.Engine), app.Config, app.logger, app.db
}
