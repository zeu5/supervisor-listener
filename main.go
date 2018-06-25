package main

import (
	"fmt"
)

func main() {
	flag1 := NewFlag("config", "Config", "bahhhh")
	s1One := NewFlag("oneFlag", "One's flag", "")
	s2One := NewFlag("twoFlag", "Two's flag", "")
	s1 := NewSubCommand("one", "First subcommand", s1One)
	s2 := NewSubCommand("two", "Second subcommand", s2One)

	AddGlobalFlag(flag1)
	AddSubCommand(s1)
	AddSubCommand(s2)

	global, sub := ParseFlags()
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
