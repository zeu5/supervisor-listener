package config

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

type Config struct {
	Handlers  map[string]HandlerConfig
	Listeners []ListenerConfig
}

func loadConfigFile(configfilepath string) (*ini.File, error) {
	var configfile *ini.File
	var err error

	configfilepaths := []string{configfilepath, "/etc/supervisor/listener.conf", "/supervisor/listener.conf"}
	filefound := false

	for _, path := range configfilepaths {
		configfile, err = ini.Load(path)
		if err != nil {
			continue
		}
		filefound = true

		break
	}
	if !filefound {
		return configfile, fmt.Errorf("Could not parse config")
	}
	return configfile, nil
}

func parseConfigFile(configfile *ini.File) *Config {
	handlerConfigs := make(map[string]HandlerConfig)
	listenerConfigs := make([]ListenerConfig, 0)

	for _, section := range configfile.Sections() {
		if strings.HasPrefix(section.Name(), "handler:") {
			handlerconfig := parseHandlerSection(section)
			handlerConfigs[handlerconfig.Name] = handlerconfig
		}
		if strings.HasPrefix(section.Name(), "listener:") {
			listenerConfigs = append(listenerConfigs, parseListenerSection(section))
		}
	}

	return &Config{
		Handlers:  handlerConfigs,
		Listeners: listenerConfigs,
	}
}

func validateConfig(config *Config) error {
	// Need to validate if the handlers referenced in the listeners are right
	// A place to add more constraints in the config if necessary in the future
	return nil
}

func ParseConfig(flags map[string]string) (*Config, error) {

	var config *Config

	configfile, err := loadConfigFile(flags["config"])
	if err != nil {
		return config, err
	}

	config = parseConfigFile(configfile)
	if err := validateConfig(config); err != nil {
		return config, err
	}
	return config, nil
}
