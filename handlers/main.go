package handlers

import (
	"github.com/zeu5/supervisor-listener/events"
)

// Handler to a supervisor event
type Handler interface {
	Init([]HandlerParam) error
	Run(events.SupervisorEvent) error
}

type HandlerParam struct {
	Name     string
	Desc     string
	Default  string
	Value    string
	Required bool
}
