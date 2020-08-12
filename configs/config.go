package configs

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	LoadApplication()
}
