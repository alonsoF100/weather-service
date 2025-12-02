package config

import (
	"log"

	"github.com/spf13/viper"
)

func Load() *Config {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	log.Printf("Config loaded successfully")
	return &config
}
