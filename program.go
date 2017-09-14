package srgnt

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/fatih/color"
)

type Program struct {
	Name     string
	Args     []string
	Commands map[string]Command
}

func (p *Program) Run() {
	flag.Parse()
	cmd := flag.Arg(0)

	if cmd == "" || cmd == "help" {
		Help(p)
	} else if val, ok := p.Commands[cmd]; ok {
		val.Callback(flag.CommandLine)
	} else {
		fmt.Printf(color.RedString("Command \"%s\" does not exist.\n"), cmd)
	}
}

func (p *Program) AddCommand(name string, callback CommandFunction, desc string) Command {
	if len(p.Commands) == 0 {
		p.Commands = make(map[string]Command)
	}
	p.Commands[name] = Command{Callback: callback, Description: desc}
	return p.Commands[name]
}

func (p *Program) AddBoolFlag(name string, value bool, usage string) {
	flag.Bool(name, value, usage)
}

func (p *Program) AddIntFlag(name string, value int, usage string) {
	flag.Int(name, value, usage)
}

func (p *Program) AddStringFlag(name string, value string, usage string) {
	flag.String(name, value, usage)
}

func Help(p *Program) {
	var b bytes.Buffer

	for k, v := range p.Commands {
		fmt.Fprintf(&b, "\t%s\t%s\n", k, v.Description)
	}

	fmt.Println(color.YellowString("Usage:"))
	fmt.Println("")
	fmt.Printf("\t%s [flags] <command>\n\n", color.CyanString(p.Name))
	fmt.Println(color.YellowString("Commands:\n"))
	fmt.Println(b.String())
}
