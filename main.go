package main

import (
	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/handlers"
)

func init() {
	setupFlags()
}

func main() {
	flags := parseFlags()
	config := config.ParseConfig(flags)
	handlers.InitHandlers(config)
	initListener(config)
	runListener()
}
