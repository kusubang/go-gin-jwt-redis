package main

import (
	"github.com/kusubang/auth/internal/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	log := logger.GetLogger()

	log.SetReportCaller(true)

	mainLogger := log.WithFields(logrus.Fields{
		"category":  "instance",
		"component": "main",
	})

	mainLogger.Info("service started")
	log.Warn("This is a warning")
	log.Error("An error occured!")

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	// r := routes.SetupRouter()
	// r.Run(":" + viper.GetString("PORT"))
}
