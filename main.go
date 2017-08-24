package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/syedomair/weather/app"
)

func main() {

	fs := flag.NewFlagSet("command", flag.ExitOnError)

	var (
		config = fs.String("config", "config/config_test.yml", "configuration file path")
	)

	/*
		        Removed for Heroku
			if len(os.Args) < 2 {
				fmt.Println("command subcommand is required")
				fs.PrintDefaults()
				os.Exit(1)
			}
	*/

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "time", log.DefaultTimestamp)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	a := app.CreateGinApplication(gin.ReleaseMode, *config, logger)
	go a.Run()

	//logger.Log("transport", "HTTP", "addr", a.Config.HttpAddress)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
