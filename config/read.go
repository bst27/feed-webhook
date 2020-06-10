package config

import (
	"encoding/json"
	"io/ioutil"
)

func Read(configFile string) (*Config, error) {
	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = json.Unmarshal(f, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
