package srgnt

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
)

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
