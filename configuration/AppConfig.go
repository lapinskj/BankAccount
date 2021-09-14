package configuration

import (
	"io/ioutil"
	"os"
	"encoding/json"

	"github.com/pkg/errors"
)
type AppConfig struct {
	Port string
	DbConfig DbConfig
}

func LoadConfig(configPath string) (AppConfig, error) {
	fName := "configuration.LoadConfig"

	f, err := os.Open(configPath)
	if err != nil {
		return AppConfig{}, errors.Wrap(err, fName)
	}
	defer f.Close()

	confBytes, readErr := ioutil.ReadAll(f)
	if readErr != nil {
		return AppConfig{}, errors.Wrap(err, fName)
	}

	var conf AppConfig
	err = json.Unmarshal(confBytes, &conf)
	if err != nil {
		return AppConfig{}, errors.Wrap(err, fName)
	}

	return conf, nil
}
