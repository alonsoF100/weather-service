package config

import "time"

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Client    ClientConfig    `mapstructure:"client"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Logger    LoggerConfig    `mapstructure:"logger"`
	Migration MigrationConfig `mapstructure:"migration"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type ClientConfig struct {
	Timeout time.Duration `mapstructure:"timeout"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSlMode  string `mapstructure:"ssl_mode"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

type MigrationConfig struct {
	Dir string `mapstructure:"dir"`
}
