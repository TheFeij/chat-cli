package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewAppConfig),
)

type Kafka struct {
	Broker  string `mapstructure:"broker"`
	Topic   string `mapstructure:"topic"`
	GroupId string `mapstructure:"group_id"`
}

type Mongo struct {
	ConnectionString string `mapstructure:"connection_string"`
	Collection       string `mapstructure:"collection"`
}

type AppConfig struct {
	Environment string `mapstructure:"environment"`
	Address     string `mapstructure:"address"`
	Port        string `mapstructure:"port"`
	Kafka       Kafka  `mapstructure:"kafka"`
	Mongo       Mongo  `mapstructure:"mongo"`
}

func NewAppConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	return &appConfig, nil
}
