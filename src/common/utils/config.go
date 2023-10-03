package utils

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Db struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Address  string `json:"address"`
		Database string `json:"database"`
	} `json:"db"`
}

func ReadConfig(filePath string) (Configuration, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Configuration{}, err
	}
	defer file.Close()

	var config Configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Configuration{}, err
	}

	return config, nil
}
