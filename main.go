package main

import (
	"github.com/kusubang/auth/internal/logger"
	"github.com/kusubang/auth/internal/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logger := logger.GetLogger()

	// log.SetReportCaller(true)

	logger.WithFields(logrus.Fields{
		"category":  "instance",
		"component": "main",
	}).Info("service started")

	logger.Warn("This is a warning")
	logger.Error("An error occured!")

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	r := routes.SetupRouter()
	r.Run(":" + viper.GetString("PORT"))
}
