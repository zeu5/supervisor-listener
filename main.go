package main

import (
	"fmt"

	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/handlers"
)

func main() {
	flags := parseFlags()
	config, err := config.ParseConfig(flags)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	handlers.InitHandlers(config)
	initListener(config)
	runListener()
}
