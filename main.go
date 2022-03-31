package main

import (
	"github.com/kusubang/auth/internal/logger"
	"github.com/kusubang/auth/internal/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var log *logrus.Logger

func main() {

	log = logger.GetLogger()

	// log.SetReportCaller(true)

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	r := routes.SetupRouter()
	r.Run(":" + viper.GetString("PORT"))
}
