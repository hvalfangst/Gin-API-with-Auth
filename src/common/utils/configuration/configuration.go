package configuration

import (
	"encoding/json"
	"errors"
	"os"
)

// Configuration represents the JSON structure.
type Configuration struct {
	Db  Db  `json:"db"`
	Jwt Jwt `json:"jwt"`
}

type Db struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Database string `json:"database"`
}

type Jwt struct {
	EncryptionKey string `json:"encryption_key"`
}

func Get(key string) (interface{}, error) {
	file, err := os.Open("src/configuration.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	// Determine which key to access and return the corresponding value
	switch key {
	case "db":
		return config.Db, nil
	case "jwt":
		return config.Jwt, nil
	default:
		return nil, errors.New("invalid key")
	}
}
