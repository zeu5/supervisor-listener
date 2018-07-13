package handlers

import (
	"github.com/zeu5/supervisor-listener/events"
)

// Handler to a supervisor event
type Handler interface {
	Run(events.SupervisorEvent) error
	SetEvent(name string)
	GetEvent() string
}

type HandlerParam struct {
	Name     string
	Desc     string
	Default  string
	Value    string
	Required bool
}
