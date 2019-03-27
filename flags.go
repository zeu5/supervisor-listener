package main

import "flag"

func parseFlags() (string, bool) {
	var configpath string
	var verbose bool
	flag.StringVar(&configpath, "config", "/etc/supervisord/listener.conf", "Path to config file")
	flag.BoolVar(&verbose, "verbose", false, "Verbosity level for logs")
	flag.Parse()

	return configpath, verbose
}
