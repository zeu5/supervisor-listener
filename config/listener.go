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

func parseListenerSection(section *ini.Section) (ListenerConfig, bool) {
	requiredkeys := []string{"events", "handlers"}
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

	otherprops := make(map[string]string)

	for _, key := range section.Keys() {
		name := key.Name()
		if name != "events" && name != "handlers" {
			otherprops[name] = key.Value()
		}
	}

	// If listening for any process related events, then process should be specified
	eventsstring := section.Key("events").Value()
	if strings.Contains(eventsstring, "PROCESS_STATE") || strings.Contains(eventsstring, "PROCESS_LOG") || strings.Contains(eventsstring, "PROCESS_COMMUNICATION") {
		if _, ok := otherprops["process"]; !ok {
			return listenerconfig, false
		}
	}

	return ListenerConfig{
		Name:     strings.Split(section.Name(), ":")[1],
		Events:   section.Key("events").Strings(","),
		Handlers: section.Key("handlers").Strings(","),
		Props:    otherprops,
	}, true
}
