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

func addProcessHandlerInstance(eventtype string, phandler Handler) {
	if _, ok := pHandlerInstances[eventtype]; !ok {
		pHandlerInstances[eventtype] = make(map[string][]Handler)
	}
	if _, ok := pHandlerInstances[eventtype][phandler.Process()]; !ok {
		pHandlerInstances[eventtype][phandler.Process()] = make([]Handler, 0)
	}
	pHandlerInstances[eventtype][phandler.Process()] = append(pHandlerInstances[eventtype][phandler.Process()], phandler)
}

func GetHandlerInstances(event *events.Event) ([]Handler, error) {
	if strings.Contains(event.Type, "PROCESS_STATE") || strings.Contains(event.Type, "PROCESS_COMMUNICATION") {
		processhandlers, ok := pHandlerInstances[event.Type]
		if !ok {
			return nil, fmt.Errorf("No handler for the eventtype")
		}
		fhandlers, ok := processhandlers[event.Body["processname"]]
		if !ok {
			return nil, fmt.Errorf("No handlers for the given process name")
		}
		return fhandlers, nil
	} else {
		handlers, ok := handlerInstances[event.Type]
		if !ok {
			return nil, fmt.Errorf("No handler for the eventtype")
		}
		return handlers, nil
	}
}
