package handlers

import (
	"github.com/zeu5/supervisor-listener/events"
)

type SlackHandler struct {
	process string
}

func (s *SlackHandler) HandleEvent(event *events.Event) error {
	return nil
}

func (s *SlackHandler) IsProcessSpecific() bool {
	return false
}

func (s *SlackHandler) Process() string {
	return s.process
}

func NewSlackHandler(props map[string]string) (Handler, error) {
	return &SlackHandler{}, nil
}
