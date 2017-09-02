package srgnt

import (
	"bytes"
	"fmt"
)

type Program struct {
	Name     string
	Args     []string
	Commands map[string]Command
}

func (p *Program) Run() {
	if len(p.Args) == 1 {
		Help(p)
	} else {
		cmd := p.Args[1]
		if val, ok := p.Commands[cmd]; ok {
			val.Callback()
		} else {
			fmt.Printf("Command \"%s\" does not exist.\n", cmd)
		}
	}
}

func (p *Program) AddCommand(name string, callback CommandFunction, desc string) {
	if len(p.Commands) == 0 {
		p.Commands = make(map[string]Command)
	}
	p.Commands[name] = Command{Callback: callback, Description: desc}
}

func Help(p *Program) {
	var b bytes.Buffer

	for k, v := range p.Commands {
		b.WriteString(fmt.Sprintf("\t%s\t%s\n", k, v.Description))
	}

	help := `
Usage: %s <command>

Commands:
%s
`
	fmt.Printf(help, p.Name, b.String())
}
