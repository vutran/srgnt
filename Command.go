package srgnt

import (
	"flag"
)

type CommandFunction func(flags *flag.FlagSet)

type Command struct {
	Description string
	Callback    CommandFunction
}

func (cmd *Command) AddBoolFlag(name string, value bool, usage string) {
	flag.Bool(name, value, usage)
}

func (cmd *Command) AddIntFlag(name string, value int, usage string) {
	flag.Int(name, value, usage)
}

func (cmd *Command) AddStringFlag(name string, value string, usage string) {
	flag.String(name, value, usage)
}
