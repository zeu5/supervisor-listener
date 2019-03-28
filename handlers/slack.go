package handlers

import (
	"github.com/zeu5/supervisor-listener/events"
)

// SlackHandler dispatches messages to slack
type SlackHandler struct {
	process string
}

// HandleEvent dispatches the event to the respective slack channel
func (s *SlackHandler) HandleEvent(event *events.Event, props map[string]string) error {
	return nil
}

// IsProcessSpecific is true if the event is specific to a process false otherwise
func (s *SlackHandler) IsProcessSpecific() bool {
	return false
}

// Process - If the event is specific to a process it returns the process name. An empty string otherwise
func (s *SlackHandler) Process() string {
	return s.process
}

// NewSlackHandler creates a new handler based on the properties provided
func NewSlackHandler(props map[string]string) (Handler, error) {
	return &SlackHandler{}, nil
}
