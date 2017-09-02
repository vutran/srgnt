package srgnt

import (
	"bytes"
	"fmt"
	"os"
)

type CommandFunction func()

type Command struct {
	Callback CommandFunction
}

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

func (p *Program) AddCommand(name string, callback CommandFunction) {
	if len(p.Commands) == 0 {
		p.Commands = make(map[string]Command)
	}
	p.Commands[name] = Command{Callback: callback}
}

func Help(p *Program) {
	var b bytes.Buffer

	for k := range p.Commands {
		b.WriteString("\t- ")
		b.WriteString(k)
		b.WriteString("\n")
	}

	help := `
Usage: %s <command>

Commands:
%s
`
	fmt.Printf(help, p.Name, b.String())
}

func CreateProgram(name string) Program {
	args := os.Args
	cli := Program{Name: name, Args: args}
	return cli
}
