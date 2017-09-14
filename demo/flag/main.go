package main

import (
	"flag"
	"fmt"
	"github.com/vutran/srgnt"
)

func Hello(flags *flag.FlagSet) {
	name := flags.Lookup("name")
	fmt.Printf("Hello, %s!\n", name.Value.String())
}

func main() {
	cli := srgnt.CreateProgram("foo")
	cli.AddCommand("hello", Hello, "Prints \"Hello, <name>!\"")
	cli.AddStringFlag("name", "", "Set a name")
	cli.Run()
}
