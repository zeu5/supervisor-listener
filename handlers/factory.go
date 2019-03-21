package handlers

var (
	handlers = make(map[string]HandlerConcstructor)
)

func registerHandler(name string, creator HandlerConcstructor) {
	handlers[name] = creator
}

func getHandler(name string) Handler {
	c := handlers[name]
	return c()
}
