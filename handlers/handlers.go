package handlers

import (
	"fmt"
	"strings"

	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
)

// Handler interface abstracts a event handler
// The event recorded from supervisor is dispatched to the handler which can then choose to handler it as necessary
type Handler interface {
	// HandlerEvent is the main method which handles the event that is dispatched form supervisor
	HandleEvent(event *events.Event, props map[string]string) error
}

// HandlerConstructor is a function which instantiates an Handler based on the given properties
type HandlerConstructor = func(map[string]string) (Handler, error)

// InitHandlers creates handler instances based on the provided configuration
func InitHandlers(config *config.Config) error {
	for _, listenerconfig := range config.Listeners {
		for _, handlername := range listenerconfig.Handlers {
			handlerconfig := config.Handlers[handlername]
			c, err := getHandlerConstructor(handlerconfig.Type)
			if err != nil {
				return err
			}
			props := handlerconfig.Props
			for k, v := range listenerconfig.Props {
				props[k] = v
			}
			h, err := c(props)
			if err != nil {
				return fmt.Errorf("Error instantiating handler instance of type %s for the listener section %s : %s", handlerconfig.Type, listenerconfig.Name, err)
			}
			for _, eventtype := range listenerconfig.Events {
				if strings.Contains(eventtype, "PROCESS") {
					addProcessHandlerInstance(eventtype, listenerconfig.Props["process"], h)
				} else {
					addHandlerInstance(eventtype, h)
				}
			}
		}
	}
	return nil
}

func init() {
	registerHandler("slack", NewSlackHandler)
}
