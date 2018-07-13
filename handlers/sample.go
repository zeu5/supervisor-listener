package handlers

import (
	"github.com/zeu5/supervisor-listener/events"
)

var (
	TestHandlerParams = []HandlerParam{
		HandlerParam{
			Name:     "one",
			Desc:     "One",
			Default:  "bah",
			Required: true,
		},
	}
)

type TestHandler struct {
	event string
}

func (t *TestHandler) Run(event events.SupervisorEvent) error {
	return nil
}

func (t *TestHandler) SetEvent(event string) {
	t.event = event
}

func (t *TestHandler) GetEvent() string {
	return t.event
}

func NewTestHandler(event string, flags []HandlerParam) Handler {
	return &TestHandler{
		event: event,
	}
}
