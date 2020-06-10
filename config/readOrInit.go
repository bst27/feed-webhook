package config

import (
	"errors"
	"os"
)

func ReadOrInit(configFile string) (*Config, error) {
	conf, err := Read(configFile)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}

		conf = initConfig()
	}

	return conf, nil
}
