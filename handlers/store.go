package handlers

var (
	handlerStore  map[string]HandlerConstructor
	handlerParams map[string][]HandlerParam
)

type HandlerConstructor func(string, []HandlerParam) Handler

func RegisterHandler(
	name string,
	handlerConstructor HandlerConstructor,
	params []HandlerParam,
) {
	if _, exists := handlerStore[name]; !exists {
		handlerStore[name] = handlerConstructor
		handlerParams[name] = params
	}
}

func Exists(name string) bool {
	_, exists := handlerStore[name]
	return exists
}

func GetAllHandlers() []string {
	var names []string
	for name := range handlerStore {
		names = append(names, name)
	}
	return names
}

func NewHandler(name string, event string, options []HandlerParam) Handler {
	if constructor, exists := handlerStore[name]; exists {
		return constructor(event, options)
	}
	return nil
}

func GetHandlerOptions(name string) []HandlerParam {
	if params, exists := handlerParams[name]; exists {
		return params
	}
	return nil
}

func GetAllHandlerOptions() map[string][]HandlerParam {
	return handlerParams
}

func init() {
	RegisterHandler("test", NewTestHandler, TestHandlerParams)
}
