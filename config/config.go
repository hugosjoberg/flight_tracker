package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       int    `mapstructure:"PORT"`
	DBPort     int    `mapstructure:"POSTGRES_PORT"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBPHOST    string `mapstructure:"POSTGRES_HOST"`
	Database   string `mapstructure:"POSTGRES_DB"`
}

var config Config

func Init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error reading config configuration file", err)
	}
	viper.AutomaticEnv()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("error when unmarshal config", err)
	}
}

func GetConfig() Config {
	return config
}
