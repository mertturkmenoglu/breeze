package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Bootstrap() {
	// First load the .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatal("Cannot read configuration file: " + err.Error())
	}
}
