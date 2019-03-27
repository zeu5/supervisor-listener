package handlers

import (
	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
)

type Handler interface {
	HandleEvent(*events.Event) error
	IsProcessSpecific() bool
	Process() string
}

type HandlerConcstructor = func(map[string]string) (Handler, error)

func InitHandlers(config *config.Config) {

}

func init() {
	registerHandler("slack", NewSlackHandler)
}
