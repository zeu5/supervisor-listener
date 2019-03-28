package handlers

import "fmt"

var (
	handlers = make(map[string]HandlerConstructor)
)

func registerHandler(name string, creator HandlerConstructor) bool {
	if _, ok := handlers[name]; ok {
		return false
	}
	handlers[name] = creator
	return true
}

func getHandlerConstructor(name string) (HandlerConstructor, error) {
	c, ok := handlers[name]
	if !ok {
		return nil, fmt.Errorf("No handler of type %s exists", name)
	}
	return c, nil
}
