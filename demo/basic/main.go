package main

import (
	"flag"
	"fmt"
	"github.com/vutran/srgnt"
)

func Hello(_ *flag.FlagSet) {
	fmt.Println("Hello, world!")
}

func main() {
	cli := srgnt.CreateProgram("foo")
	cli.AddCommand("hello", Hello, "Prints \"Hello, world!\"")
	cli.Run()
}
