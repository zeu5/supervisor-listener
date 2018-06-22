package options

import "supervisor-listener/handlers"

type Options struct {
	events   []string
	handlers []handlers.Handler
}

func NewOptions() *Options {
	return &Options{}
}
