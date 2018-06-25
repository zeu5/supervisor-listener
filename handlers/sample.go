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

type TestHandler struct{}

func (t *TestHandler) Init(flags []HandlerParam) error {
	return nil
}

func (t *TestHandler) Run(event events.SupervisorEvent) error {
	return nil
}

func NewTestHandler(flags []HandlerParam) Handler {
	return &TestHandler{}
}
