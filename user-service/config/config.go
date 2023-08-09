package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConfig Database `mapstructure:"database"`
	SecretConfig   Secret   `mapstructure:"secret"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Database string `mapstructure:"database"`
}

type Secret struct {
	JWT      string `mapstructure:"jwt"`
	Database string `mapstructure:"database"`
}

func InitConfig() (*Config, error) {
	viper.SetConfigFile("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Unable to decode config into struct", err)
	}
	return &config, nil
}
