package gearconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type GearConfig struct {
	Language string `json:"language"`
}

func GetConfigContent() ([]byte, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return []byte{}, err
	}
	configFilePath := dirname + "/.gear_config.json"
	if _, err := os.Stat(configFilePath); err != nil {
		return []byte{}, err
	}
	return ioutil.ReadFile(configFilePath)
}

func GetConfig() (GearConfig, error) {
	bytes, err := GetConfigContent()
	if err != nil {
		return GearConfig{}, err
	}
	var config GearConfig
	if json.Unmarshal(bytes, &config) != nil {
		return GearConfig{}, err
	}
	return config, nil
}
