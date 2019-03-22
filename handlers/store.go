package handlers

import "github.com/zeu5/supervisor-listener/events"

var (
	handlerInstances map[string]Handler
)

func addInstance(name string, handler Handler) {
	handlerInstances[name] = handler
}

func GetHandlerInstance(event *events.Event) (Handler, error) {
	return &SlackHandler{}, nil
}
