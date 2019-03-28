package handlers

import (
	"strings"

	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
)

type Handler interface {
	HandleEvent(event *events.Event, props map[string]string) error
	IsProcessSpecific() bool
	Process() string
}

type HandlerConcstructor = func(map[string]string) (Handler, error)

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
