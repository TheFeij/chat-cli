package config

import (
	"github.com/spf13/viper"
	"os"
)

type Kafka struct {
	Broker  string `mapstructure:"broker" json:"broker" yaml:"broker"`
	Topic   string `mapstructure:"topic" json:"topic" yaml:"topic"`
	GroupId string `mapstructure:"group_id" json:"group_id" yaml:"group_id"`
}

type Mongo struct {
	ConnectionString string `mapstructure:"connection_string" json:"connection_string" yaml:"connection_string"`
	Collection       string `mapstructure:"collection" json:"collection" yaml:"collection"`
}

type AppConfig struct {
	ServerAddress string `mapstructure:"server_address" json:"server_address" yaml:"server_address"`
	ServerPort    string `mapstructure:"server_port"  json:"server_port" yaml:"server_port"`
	Kafka         Kafka  `mapstructure:"kafka" json:"kafka" yaml:"kafka"`
	Mongo         Mongo  `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
}

const (
	configFilePathKey = "CONFIG_FILE_PATH"
	configFileNameKey = "CONFIG_FILE_NAME"
	configFileTypeKey = "CONFIG_FILE_TYPE"
)

func NewAppConfig() (*AppConfig, error) {
	filepath := os.Getenv(configFilePathKey)
	if filepath == "" {
		return nil, ErrEnvFilePathNotSet
	}

	filename := os.Getenv(configFileNameKey)
	if filename == "" {
		return nil, ErrEnvFileNameNotSet
	}

	filetype := os.Getenv(configFileTypeKey)
	if filetype == "" {
		return nil, ErrEnvFileTypeNotSet
	}

	viper.SetConfigName(filename)
	viper.SetConfigType(filetype)
	viper.AddConfigPath(filepath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, newErrReadingConfig(err)
	}

	var appConfig AppConfig

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, newErrUnmarshallingConfig(err)
	}

	return &appConfig, nil
}
