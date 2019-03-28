package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/handlers"
)

func main() {

	configpath, verbose := parseFlags()
	initLogger(verbose)
	config, err := config.ParseConfig(configpath)
	if err != nil {
		log.Fatal(err)
	}
	if err := handlers.InitHandlers(config); err != nil {
		log.Fatal(err)
	}

	initListener(config)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigint, syscall.SIGTERM)

	runListener(sigint)
}
