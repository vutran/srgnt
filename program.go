package srgnt

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
)

type Program struct {
	Name     string
	Args     []string
	Commands map[string]Command
}

func (p *Program) Run(done chan bool) {
	flag.Parse()
	cmd := flag.Arg(0)

	var b bytes.Buffer
	var r io.Reader
	var d bool

	if cmd == "" || cmd == "help" {
		r = Help(p)
	} else if val, ok := p.Commands[cmd]; ok {
		r = val.Callback(flag.CommandLine)
	} else {
		fmt.Fprintf(&b, color.RedString("Command \"%s\" does not exist.\n"), cmd)
	}

	if r != nil {
		b.ReadFrom(r)
	}

	fmt.Printf(b.String())

	done <- d
}

func (p *Program) AddCommand(name string, callback CommandFunction, desc string) Command {
	if len(p.Commands) == 0 {
		p.Commands = make(map[string]Command)
	}
	p.Commands[name] = Command{Callback: callback, Description: desc}
	return p.Commands[name]
}

func Help(p *Program) io.Reader {
	var ret, b bytes.Buffer

	for k, v := range p.Commands {
		fmt.Fprintf(&b, "\t%s\t%s\n", k, v.Description)
	}

	fmt.Fprintf(&ret, color.YellowString("Usage:\n"))
	fmt.Fprintf(&ret, "\n")
	fmt.Fprintf(&ret, "\t%s <command>\n\n", color.CyanString(p.Name))
	fmt.Fprintf(&ret, color.YellowString("Commands:\n\n"))
	fmt.Fprintf(&ret, "%s", b.Bytes())

	return &ret
}