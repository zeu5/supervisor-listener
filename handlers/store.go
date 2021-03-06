package handlers

import (
	"fmt"
	"strings"

	"github.com/zeu5/supervisor-listener/events"
)

var (
	handlerInstances  = make(map[string][]Handler)
	pHandlerInstances = make(map[string]map[string][]Handler)
)

func addHandlerInstance(eventtype string, handler Handler) {
	if _, ok := handlerInstances[eventtype]; !ok {
		handlerInstances[eventtype] = make([]Handler, 0)
	}
	handlerInstances[eventtype] = append(handlerInstances[eventtype], handler)
}

func addProcessHandlerInstance(eventtype, process string, phandler Handler) {
	if _, ok := pHandlerInstances[eventtype]; !ok {
		pHandlerInstances[eventtype] = make(map[string][]Handler)
	}
	if _, ok := pHandlerInstances[eventtype][process]; !ok {
		pHandlerInstances[eventtype][process] = make([]Handler, 0)
	}
	pHandlerInstances[eventtype][process] = append(pHandlerInstances[eventtype][process], phandler)
}

// GetHandlerInstances returns the instances which concern the specific event
// The handler instances are organised based on the event type and process type.
// If the event is specific to a process then the respective handlers are looked up in pHandlerInstances. handlerInstances otherwise.
func GetHandlerInstances(event *events.Event) ([]Handler, error) {
	if strings.Contains(event.Type, "PROCESS_STATE") || strings.Contains(event.Type, "PROCESS_COMMUNICATION") {
		processhandlers, ok := pHandlerInstances[event.Type]
		if !ok {
			return nil, fmt.Errorf("No handler for the eventtype : %s", event.Type)
		}
		handlers, ok := processhandlers[event.Body["processname"]]
		if !ok {
			return nil, fmt.Errorf("No handlers for the given process name : %s", event.Body["processname"])
		}
		return handlers, nil
	}

	handlers, ok := handlerInstances[event.Type]
	if !ok {
		return nil, fmt.Errorf("No handler for the eventtype : %s", event.Type)
	}
	return handlers, nil
}
