package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SERVER_PORT       int
	DEBUG_SERVER_PORT int
	DB_PORT           int
	DB_USERNAME       string
	DB_PASSWORD       string
	DB_HOST           string
	DB_NAME           string
}

func LoadConfig(path string) (config Config) {
	viper.SetConfigFile(fmt.Sprintf("%s.env", path))
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Cannot load the config file")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Cannot load the config file")
	}
	return
}
