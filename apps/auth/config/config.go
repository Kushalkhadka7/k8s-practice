package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server *serverConfig
}

type serverConfig struct {
	Port         int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

func New() (*Config, error) {
	var c Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &Config{
		Server: &serverConfig{
			Port:         c.Server.Port,
			ReadTimeOut:  c.Server.ReadTimeOut * time.Second,
			WriteTimeOut: c.Server.WriteTimeOut * time.Second,
		},
	}, nil

}
