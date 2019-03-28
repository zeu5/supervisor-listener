package handlers

import "fmt"

var (
	handlers = make(map[string]HandlerConcstructor)
)

func registerHandler(name string, creator HandlerConcstructor) bool {
	if _, ok := handlers[name]; ok {
		return false
	}
	handlers[name] = creator
	return true
}

func getHandlerConstructor(name string) (HandlerConcstructor, error) {
	c, ok := handlers[name]
	if !ok {
		return nil, fmt.Errorf("No handler of type %s exists", name)
	}
	return c, nil
}
