package handlers

var (
	handlers map[string]HandlerConcstructor
)

func registerHandler(name string, creator HandlerConcstructor) {
	handlers[name] = creator
}

func getHandler(name string) Handler {
	c := handlers[name]
	return c()
}
