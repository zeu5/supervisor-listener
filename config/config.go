package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

// Config is the structure to hold the parsed configuration from the ini config file
type Config struct {
	GlobalProps map[string]string
	Handlers    map[string]HandlerConfig
	Listeners   []ListenerConfig
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
		return configfile, fmt.Errorf("Config file not found")
	}
	return configfile, nil
}

func parseConfigFile(configfile *ini.File) *Config {
	handlerConfigs := make(map[string]HandlerConfig)
	listenerConfigs := make([]ListenerConfig, 0)
	globalprops := make(map[string]string)

	defaultsection := configfile.Section(ini.DEFAULT_SECTION)
	if defaultsection.HasKey("devicename") {
		globalprops["devicename"] = defaultsection.Key("devicename").String()
	} else {
		if hostname, err := os.Hostname(); err == nil {
			globalprops["devicename"] = hostname
		} else {
			globalprops["devicename"] = ""
		}
	}

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
		Handlers:    handlerConfigs,
		Listeners:   listenerConfigs,
		GlobalProps: globalprops,
	}
}

func validateConfig(config *Config) error {
	// A place to add more constraints in the config if necessary in the future

	// Handler section specific checks
	for _, handlerconfig := range config.Handlers {
		if handlerconfig.Type == "" {
			return fmt.Errorf("Handler section has empty type")
		}
	}

	// Listener section specific checks
	for _, listenerconfig := range config.Listeners {
		if len(listenerconfig.Events) == 0 {
			return fmt.Errorf("Listener section %s is not subscribed to any events", listenerconfig.Name)
		}
		if len(listenerconfig.Handlers) == 0 {
			return fmt.Errorf("Listener section %s does not have any handlers", listenerconfig.Name)
		}
		isProcSpecific := false
		for _, event := range listenerconfig.Events {
			if strings.Contains(event, "PROCESS") {
				isProcSpecific = true
				break
			}
		}
		if _, ok := listenerconfig.Props["process"]; isProcSpecific && !ok {
			return fmt.Errorf("Listener section %s is missing process key but has subscribed to process events", listenerconfig.Name)
		}
	}

	for _, listenerconfig := range config.Listeners {
		for _, handlername := range listenerconfig.Handlers {
			if _, ok := config.Handlers[handlername]; !ok {
				return fmt.Errorf("Handler of type %s does not exist in listener section %s", handlername, listenerconfig.Name)
			}
		}
	}
	return nil
}

// ParseConfig takes a file path and parses an ini config in that path or one of the default paths
func ParseConfig(configpath string) (*Config, error) {

	var config *Config

	configfile, err := loadConfigFile(configpath)
	if err != nil {
		return config, err
	}

	config = parseConfigFile(configfile)
	if err := validateConfig(config); err != nil {
		return config, err
	}
	return config, nil
}
