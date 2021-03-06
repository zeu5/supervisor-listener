package handlers

import (
	"fmt"

	"github.com/zeu5/supervisor-listener/events"
)

// SlackHandler dispatches messages to slack
type SlackHandler struct {
	process string
}

// HandleEvent dispatches the event to the respective slack channel
func (s *SlackHandler) HandleEvent(event *events.Event, props map[string]string) error {
	fmt.Println(event.Body, event.Type, event.Header, props)
	return nil
}

// NewSlackHandler creates a new handler based on the properties provided
func NewSlackHandler(props map[string]string) (Handler, error) {
	proc, ok := props["process"]
	if ok {
		return &SlackHandler{
			process: proc,
		}, nil
	}
	return &SlackHandler{}, nil
}
