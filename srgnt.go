package srgnt

import (
	"os"
)

func CreateProgram(name string) Program {
	args := os.Args
	cli := Program{Name: name, Args: args}
	return cli
}
