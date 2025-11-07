package config_test

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"message-persist/internal/app/infrastructure/config"
	"os"
	"testing"
)

func TsFillAppConfig(condition *config.AppConfig) {
	if condition == nil {
		condition = &config.AppConfig{}
	}

	if condition.ServerAddress == "" {
		condition.ServerAddress = "localhost"
	}
	if condition.ServerPort == "" {
		condition.ServerPort = "8080"
	}

	if condition.Kafka.Topic == "" {
		condition.Kafka.Topic = "test"
	}
	if condition.Kafka.Broker == "" {
		condition.Kafka.Broker = "localhost:9092"
	}
	if condition.Kafka.GroupId == "" {
		condition.Kafka.GroupId = "test"
	}

	if condition.Mongo.ConnectionString == "" {
		condition.Mongo.ConnectionString = "mongodb://localhost:27017"
	}
	if condition.Mongo.Collection == "" {
		condition.Mongo.Collection = "test"
	}
}

func TsCreateTestConfigFile(t *testing.T, condition *config.AppConfig) {
	file, err := os.Create(configFilePathValue + "/" + configFileNameValue + "." + configFileTypeValue)
	if err != nil {
		t.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		require.NoError(t, err)
	}(file)

	err = json.NewEncoder(file).Encode(condition)
	require.NoError(t, err)

	return
}

func TestNewAppConfig(t *testing.T) {
	expected := new(config.AppConfig)

	TsFillAppConfig(expected)

	TsCreateTestConfigFile(t, expected)

	actual, err := config.NewAppConfig()
	require.NoError(t, err)

	require.Equal(t, expected, actual)
}

func TestNewAppConfig_ErrReadingConfig(t *testing.T) {
	err := TsDeleteTestConfigFile()
	require.NoError(t, err)

	_, err = config.NewAppConfig()
	require.Error(t, err)

	require.ErrorIs(t, err, config.ErrReadingConfig)
}

func TestNewAppConfig_ErrUnmarshallingConfig(t *testing.T) {
	file, err := os.Create(configFilePathValue + "/" + configFileNameValue + "." + configFileTypeValue)
	if err != nil {
		t.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		require.NoError(t, err)
	}(file)

	_, err = file.WriteString(`{"kafka": 23}`)
	require.NoError(t, err)

	_, err = config.NewAppConfig()
	require.Error(t, err)

	require.ErrorIs(t, err, config.ErrUnmarshallingConfig)
}

func TestNewAppConfig_ErrEnvFileTypeNotSet(t *testing.T) {
	os.Clearenv()

	err := os.Setenv(configFilePathKey, configFilePathValue)
	require.NoError(t, err)

	err = os.Setenv(configFileNameKey, configFileNameValue)
	require.NoError(t, err)

	_, err = config.NewAppConfig()
	require.Error(t, err)

	require.Error(t, err, config.ErrEnvFileTypeNotSet)
}
func TestNewAppConfig_ErrEnvFilePathNotSet(t *testing.T) {
	os.Clearenv()

	err := os.Setenv(configFileNameKey, configFileNameValue)
	require.NoError(t, err)

	err = os.Setenv(configFileTypeKey, configFileTypeValue)
	require.NoError(t, err)

	_, err = config.NewAppConfig()
	require.Error(t, err)

	require.Error(t, err, config.ErrEnvFilePathNotSet)
}
func TestNewAppConfig_ErrEnvFileNameNotSet(t *testing.T) {
	os.Clearenv()

	err := os.Setenv(configFilePathKey, configFilePathValue)
	require.NoError(t, err)

	err = os.Setenv(configFileTypeKey, configFileTypeValue)
	require.NoError(t, err)

	_, err = config.NewAppConfig()
	require.Error(t, err)

	require.Error(t, err, config.ErrEnvFileNameNotSet)
}
