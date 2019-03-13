package main

import (
	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
	"github.com/zeu5/supervisor-listener/handlers"
)

func initListener(config *config.Config) {
	// Initialise log channels and other options
}

func processEvent(header map[string]string, eventstring string) {
	event := events.GetEvent(header, eventstring)
	handler := handlers.GetHandlerInstance(event)
	handler.HandleEvent(event)
}

func runListener() {
	for {
		header := readHeaderData()
		eventstring := readEventData()
		go processEvent(header, eventstring)
		replyOK()
	}
}
