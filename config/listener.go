package config

import (
	"strings"

	"gopkg.in/ini.v1"
)

type ListenerConfig struct {
	Name     string
	Events   []string
	Handlers []string
	Props    map[string]string
}

func parseListenerSection(section *ini.Section) ListenerConfig {

	otherprops := make(map[string]string)

	for _, key := range section.Keys() {
		name := key.Name()
		if name != "events" && name != "handlers" {
			otherprops[name] = key.Value()
		}
	}

	return ListenerConfig{
		Name:     strings.Split(section.Name(), ":")[1],
		Events:   section.Key("events").Strings(","),
		Handlers: section.Key("handlers").Strings(","),
		Props:    otherprops,
	}
}
