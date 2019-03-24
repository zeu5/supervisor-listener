package handlers

import (
	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
)

type Handler interface {
	HandleEvent(*events.Event) error
}
type HandlerConcstructor = func() Handler

func InitHandlers(config *config.Config) {

}

func init() {
	registerHandler("slack", NewSlackHandler)
}
