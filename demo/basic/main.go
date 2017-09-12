package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/vutran/srgnt"
	"io"
)

func Hello(_ *flag.FlagSet) io.Reader {
	var b bytes.Buffer

	fmt.Fprintln(&b, "Hello, world!")

	return &b
}

func main() {
	done := make(chan bool, 1)

	cli := srgnt.CreateProgram("foo")

	cli.AddCommand("hello", Hello, "Prints \"Hello, world!\"")

	cli.Run(done)

	<-done
}
