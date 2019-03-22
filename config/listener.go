package config

import (
	"strings"

	"gopkg.in/ini.v1"
)

type ListenerConfig struct {
	Name        string
	ProcessName string
	Events      []string
	Handlers    []string
}

func parseListenerSection(section *ini.Section) (ListenerConfig, bool) {
	requiredkeys := []string{"process", "events", "handlers"}
	haskeys := true
	for _, key := range requiredkeys {
		if !section.HasKey(key) {
			haskeys = false
			break
		}
	}

	var listenerconfig ListenerConfig
	if !haskeys {
		return listenerconfig, false
	}

	return ListenerConfig{
		Name:        strings.Split(section.Name(), ":")[1],
		ProcessName: section.Key("process").String(),
		Events:      section.Key("events").Strings(","),
		Handlers:    section.Key("handlers").Strings(","),
	}, true
}
