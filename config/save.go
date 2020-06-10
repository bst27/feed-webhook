package config

import (
	"encoding/json"
	"io/ioutil"
)

func Save(config *Config, filePath string) error {
	data, err := json.MarshalIndent(config, "", "   ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, 0664)
}
