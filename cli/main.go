package cli

import (
	"flag"
	"fmt"
)

var (
	subcommands = make(map[string]*SubCommand)
	globalflags = make(map[string]*Flag)
)

type Flag struct {
	Name    string
	Desc    string
	Default string
	Value   string
}

func NewFlag(name string, desc string, defaultValue string) *Flag {
	return &Flag{
		Name:    name,
		Desc:    desc,
		Default: defaultValue,
	}
}

type SubCommand struct {
	Name    string
	Flags   map[string]*Flag
	flagset *flag.FlagSet
	Desc    string
}

func NewSubCommand(name string, desc string, flags ...*Flag) *SubCommand {
	subcommand := &SubCommand{
		Name:    name,
		Flags:   make(map[string]*Flag),
		flagset: flag.NewFlagSet(name, flag.ExitOnError),
		Desc:    desc,
	}
	for _, flag := range flags {
		if _, exists := subcommand.Flags[flag.Name]; !exists {
			subcommand.Flags[flag.Name] = flag
		}
	}
	return subcommand
}

func (s *SubCommand) AddFlag(flag *Flag) {
	if _, exists := s.Flags[flag.Name]; !exists {
		s.Flags[flag.Name] = flag
	}
}

func (s *SubCommand) RemoveFlag(flagName string) {
	if _, exists := s.Flags[flagName]; exists {
		delete(s.Flags, flagName)
	}
}

func AddSubCommand(s *SubCommand) {
	for flagName, flagVal := range s.Flags {
		s.flagset.StringVar(&flagVal.Value, flagName, flagVal.Default, flagVal.Desc)
	}
	subcommands[s.Name] = s
}

func AddGlobalFlag(globalflag *Flag) {
	globalflags[globalflag.Name] = globalflag
	flag.StringVar(&globalflag.Value, globalflag.Name, globalflag.Default, globalflag.Desc)
}

func usageFunc() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "Subcommands: \n")
	maxLen := 0
	for name, _ := range subcommands {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}
	for name, subcommands := range subcommands {
		fmt.Fprintf(flag.CommandLine.Output(), "   %[3]*[1]s\t%[2]s\n", name, subcommands.Desc, maxLen)
	}
	fmt.Fprintf(flag.CommandLine.Output(), "Run <subcommand> --help for usage of the subcommand\n")
}

func ParseFlags() (map[string]*Flag, *SubCommand) {
	flag.Usage = usageFunc
	flag.Parse()

	var parsedsubcommand *SubCommand
	if s, exists := subcommands[flag.Arg(0)]; exists {
		s.flagset.Parse(flag.Args()[1:])
		parsedsubcommand = s
	}
	return globalflags, parsedsubcommand
}
