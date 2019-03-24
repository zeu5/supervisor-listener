package handlers

import (
	"fmt"
	"strings"

	"github.com/zeu5/supervisor-listener/events"
)

var (
	handlerInstances map[string][]Handler
)

func addHandlerInstance(eventtype string, handler Handler) {
	if _, ok := handlerInstances[eventtype]; !ok {
		handlerInstances[eventtype] = make([]Handler, 0)
	}
	handlerInstances[eventtype] = append(handlerInstances[eventtype], handler)
}

func GetHandlerInstance(event *events.Event) (Handler, error) {
	_, ok := handlerInstances[event.Type]
	if !ok {
		return nil, fmt.Errorf("No instances for the eventtype")
	}
	if strings.Contains(event.Type, "PROCESS_STATE") || strings.Contains(event.Type, "PROCESS_COMMUNICATION") {
	}
	return &SlackHandler{}, nil
}
