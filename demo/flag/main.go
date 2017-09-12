package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/vutran/srgnt"
	"io"
)

func Hello(flags *flag.FlagSet) io.Reader {
	name := flags.Lookup("name")

	var b bytes.Buffer

	fmt.Fprintf(&b, "Hello, %s!\n", name.Value.String())

	return &b
}

func main() {
	done := make(chan bool, 1)

	cli := srgnt.CreateProgram("foo")

	cli.AddCommand("hello", Hello, "Prints \"Hello, <name>!\"")
	cli.AddStringFlag("name", "", "Set a name")

	cli.Run(done)

	<-done
}
