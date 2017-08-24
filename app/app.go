package app

import (
	"io/ioutil"

	"github.com/go-kit/kit/log"
	"github.com/syedomair/weather/common"
	"github.com/syedomair/weather/models"

	"github.com/gin-gonic/gin"

	yaml "gopkg.in/yaml.v2"
)

func CreateGinApplication(releaseMode, configFilePath string, logger log.Logger) *GinApplication {
	gin.SetMode(releaseMode)
	ap := &GinApplication{}
	ap.logger = logger
	ap.loadConfig(configFilePath)
	ap.initialize()
	ap.routers()
	return ap
}

type GinApplication struct {
	Config common.Config
	db     models.Datastore
	Engine *gin.Engine
	logger log.Logger
}

func (a GinApplication) Run() {
	a.Engine.Run(a.Config.HttpAddress)
	/* heroku_branch */
	//a.Engine.Run(":" + os.Getenv("PORT"))
}

func (a *GinApplication) initialize() {
	a.db, _ = models.NewDB(a.Config.DatabaseURL)
}

func (a *GinApplication) loadConfig(filepath string) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &a.Config)
}

func (a *GinApplication) routers() {
	a.Engine = gin.Default()
	a.Engine.Use(dbSetup(a.db))

	v1 := a.Engine.Group("/v1")
	{
		v1.GET("/weather-log", func(c *gin.Context) {
			WeatherGetAction(c)
		})
		v1.POST("/weather", func(c *gin.Context) {
			WeatherPostAction(c, a.Config, a.logger)
		})
	}
}
func dbSetup(db models.Datastore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
