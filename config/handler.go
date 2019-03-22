package config

import (
	"strings"

	"gopkg.in/ini.v1"
)

type HandlerConfig struct {
	Name  string
	Type  string
	Props map[string]string
}

func parseHandlerSection(section *ini.Section) (HandlerConfig, bool) {
	// Assumes that section has a valid handler name

	var handlerconfig HandlerConfig

	if !section.HasKey("type") {
		return handlerconfig, false
	}

	handlername := strings.Split(section.Name(), ":")[1]
	handlertype := section.Key("type").String()
	handlerprops := make(map[string]string)
	for _, key := range section.Keys() {
		if key.Name() != "type" {
			handlerprops[key.Name()] = key.String()
		}
	}
	return HandlerConfig{
		Name:  handlername,
		Type:  handlertype,
		Props: handlerprops,
	}, true
}
