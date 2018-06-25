package handlers

var (
	handlerStore  map[string]HandlerConstructor
	handlerParams map[string][]HandlerParam
)

type HandlerConstructor func([]HandlerParam) Handler

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

func GetHandler(name string, options []HandlerParam) Handler {
	if constructor, exists := handlerStore[name]; exists {
		return constructor(options)
	}
	return nil
}

func GetAllHandlerOptions() map[string][]HandlerParam {
	return handlerParams
}

func init() {
	RegisterHandler("sample", NewTestHandler, TestHandlerParams)
}
