package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type HandlerConfig struct {
}
type ListenerConfig struct {
}
type Config struct {
	Name string `ini:"name"`
}

func ParseConfig(flags map[string]string) (*Config, error) {

	configfilepaths := []string{flags["config"], "/etc/supervisor/listener.conf", "/supervisor/listener.conf"}
	filefound := false

	config := new(Config)
	for _, path := range configfilepaths {
		configfile, err := ini.Load(path)
		if err != nil {
			continue
		}
		filefound = true
		configfile.Section("").MapTo(config)
		break
	}
	if !filefound {
		var cfg *Config
		return cfg, fmt.Errorf("Could not parse config")
	}

	return config, nil

}
