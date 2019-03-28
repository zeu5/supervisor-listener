package handlers

import (
	"strings"

	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
)

// Handler interface abstracts a event handler
// The event recorded from supervisor is dispatched to the handler which can then choose to handler it as necessary
type Handler interface {
	// HandlerEvent is the main method which handles the event that is dispatched form supervisor
	HandleEvent(event *events.Event, props map[string]string) error
	// IsProcessSpecific determines if the handler instance listens to a process specific event or not
	IsProcessSpecific() bool
	// Process - If the event is specific to a process it returns the process name. An empty string otherwise
	Process() string
}

// HandlerConstructor is a function which instantiates an Handler based on the given properties
type HandlerConstructor = func(map[string]string) (Handler, error)

// InitHandlers creates handler instances based on the provided configuration
func InitHandlers(config *config.Config) error {
	for _, listenerconfig := range config.Listeners {
		for _, handlertype := range listenerconfig.Handlers {
			handlerconfig := config.Handlers[handlertype]
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
				return err
			}
			for _, eventtype := range listenerconfig.Events {
				if strings.Contains(eventtype, "PROCESS") {
					addProcessHandlerInstance(eventtype, h)
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
