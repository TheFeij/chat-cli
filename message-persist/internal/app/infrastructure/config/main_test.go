package config_test

import (
	"log"
	"os"
	"testing"
)

const (
	configFilePathKey = "CONFIG_FILE_PATH"
	configFileNameKey = "CONFIG_FILE_NAME"
	configFileTypeKey = "CONFIG_FILE_TYPE"
)

const (
	configFilePathValue = "."
	configFileNameValue = "config.test"
	configFileTypeValue = "json"
)

func TestMain(m *testing.M) {
	TsSetEnvs(m)

	code := m.Run()

	err := TsDeleteTestConfigFile()
	if err != nil {
		log.Printf("Error deleting test config file: %s", err.Error())
	}

	os.Exit(code)
}

func TsSetEnvs(m *testing.M) {
	err := os.Setenv(configFilePathKey, configFilePathValue)
	if err != nil {
		panic(err)
	}

	err = os.Setenv(configFileNameKey, configFileNameValue)
	if err != nil {
		panic(err)
	}

	err = os.Setenv(configFileTypeKey, configFileTypeValue)
	if err != nil {
		panic(err)
	}
}

func TsDeleteTestConfigFile() error {
	err := os.Remove(configFilePathValue + "/" + configFileNameValue + "." + configFileTypeValue)
	if err == nil {
		return nil
	}

	if isNotExist := os.IsNotExist(err); isNotExist {
		return nil
	}

	return err
}
