package config

import (
	"strings"

	"gopkg.in/ini.v1"
)

type ListenerConfig struct {
	Name        string
	ProcessName string
	Events      []string
	Handler     string
}

func parseListenerSection(section *ini.Section) ListenerConfig {

	return ListenerConfig{
		Name:        strings.Split(section.Name(), ":")[1],
		ProcessName: section.Key("process").String(),
		Events:      section.Key("events").Strings(","),
		Handler:     section.Key("handler").String(),
	}
}
