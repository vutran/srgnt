package srgnt

import (
	"bytes"
	"flag"
	"fmt"
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
		fmt.Printf("Command \"%s\" does not exist.\n", cmd)
	}
}

func (p *Program) AddCommand(name string, callback CommandFunction, desc string) Command {
	if len(p.Commands) == 0 {
		p.Commands = make(map[string]Command)
	}
	p.Commands[name] = Command{Callback: callback, Description: desc}
	return p.Commands[name]
}

func Help(p *Program) {
	var b bytes.Buffer

	for k, v := range p.Commands {
		b.WriteString(fmt.Sprintf("\t%s\t%s\n", k, v.Description))
	}

	help := `
Usage:

	%s <command>

Commands:

%s
`
	fmt.Printf(help, p.Name, b.String())
}
