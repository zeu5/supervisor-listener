package main

import (
	"fmt"

	"github.com/zeu5/supervisor-listener/cli"
	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/errors"
	"github.com/zeu5/supervisor-listener/handlers"
)

func setupCli() {
	cli.AddGlobalFlag(
		cli.NewFlag("config", "Path to config file", "/etc/supervisor/listener.yml"),
	)

	handlers := handlers.GetAllHandlerOptions()
	for name, handleroptions := range handlers {

		var flags []*cli.Flag
		for _, option := range handleroptions {
			flags = append(flags, cli.NewFlag(option.Name, option.Desc, option.Default))
		}

		cli.NewSubCommand(
			name,
			fmt.Sprintf("Command for %s handler", name),
			flags...,
		)
	}
}

func getHandler(name string, flags map[string]*cli.Flag) (handlers.Handler, error) {
	options := handlers.GetHandlerOptions(name)
	var params []handlers.HandlerParam
	var event string

	if _, ok := flags["event"]; ok {
		event = flags["event"].Value
	} else {
		event = "ALL"
	}

	for _, option := range options {
		flag := flags[option.Name]
		if flag.Value == option.Default && option.Required {
			return nil, errors.MissingRequiredCommandFlag(name, flag.Name)
		}
		option.Value = flag.Value
		params = append(params, option)
	}
	handler := handlers.NewHandler(name, event, params)
	return handler, nil
}

func main() {
	setupCli()
	globalflag, subcommand := cli.ParseFlags()
	_, err := config.ParseConfig(globalflag["config"].Value)
	if err == nil {
		fmt.Println("Successfully parse config")
	} else if subcommand != nil {
		_, err = getHandler(subcommand.Name, subcommand.Flags)
	}

	if err != nil {
		fmt.Println("Not able to start")
	}
}
