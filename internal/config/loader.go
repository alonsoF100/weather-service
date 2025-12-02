package config

import (
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

	return &config
}
