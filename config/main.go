package config

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/zeu5/supervisor-listener/errors"
	"github.com/zeu5/supervisor-listener/handlers"

	yaml "gopkg.in/yaml.v2"
)

type handlerConfig struct {
	name   string
	config map[string]string
}

type eventConfig struct {
	name     string
	handlers []handlerConfig
}

type Config struct {
	handlers         []handlerConfig
	events           []eventConfig
	handlerconfigmap map[string]interface{}
}

func (c *Config) CreateHandlers() ([]handlers.Handler, error) {
	var outHandlers []handlers.Handler
	for _, handler := range c.handlers {
		if _, exists := c.handlerconfigmap[handler.name]; exists {
			return outHandlers, errors.DuplicateHandlerConfig(handler.name)
		}
		c.handlerconfigmap[handler.name] = handler.config
	}

	for _, event := range c.events {
		fmt.Println(event.name)
	}
	return outHandlers, nil
}

func ParseConfig(file string) (Config, error) {
	file = path.Clean(file)
	data, err := ioutil.ReadFile(file)

	config := Config{}

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
