package main

import (
	"fmt"

	"github.com/zeu5/supervisor-listener/cli"
)

func main() {
	flag1 := cli.NewFlag("config", "Config", "bahhhh")
	s1One := cli.NewFlag("oneFlag", "One's flag", "")
	s2One := cli.NewFlag("twoFlag", "Two's flag", "")
	s1 := cli.NewSubCommand("one", "First subcommand", s1One)
	s2 := cli.NewSubCommand("two", "Second subcommand", s2One)

	cli.AddGlobalFlag(flag1)
	cli.AddSubCommand(s1)
	cli.AddSubCommand(s2)

	global, sub := cli.ParseFlags()
	for _, f := range global {
		fmt.Println(f.Name)
		fmt.Println(f.Value)
	}
	if sub != nil {
		for _, f := range sub.Flags {
			fmt.Println(f.Name)
			fmt.Println(f.Value)
		}
	}
}
