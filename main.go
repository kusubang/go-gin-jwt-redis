package main

import (
	"github.com/kusubang/auth/internal/routes"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	r := routes.SetupRouter()
	r.Run(":" + viper.GetString("PORT"))
}
