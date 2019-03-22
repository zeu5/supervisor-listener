package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigint, syscall.SIGTERM)

	runListener(sigint)
}
