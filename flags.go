package main

import "flag"

func parseFlags() map[string]string {
	var configpath string
	flag.StringVar(&configpath, "config", "/etc/supervisord/listener.conf", "Path to config file")
	flag.Parse()

	flags := make(map[string]string)
	flags["config"] = configpath

	return flags
}
