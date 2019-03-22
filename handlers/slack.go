package handlers

import (
	"github.com/zeu5/supervisor-listener/events"
)

type SlackHandler struct {
}

func (s *SlackHandler) HandleEvent(event *events.Event) error {
	return nil
}

func NewSlackHandler() Handler {
	return &SlackHandler{}
}
