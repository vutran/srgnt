package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/vutran/srgnt"
	"io"
)

// TODO - return a map of flags instead (key ,value)
func Foo(flags *flag.FlagSet) io.Reader {
	name := flags.Lookup("name")

	var b bytes.Buffer

	fmt.Fprintf(&b, "Foo: %s\n", name.Value.String())

	return &b
}

func Bar(flags *flag.FlagSet) io.Reader {
	var b bytes.Buffer

	fmt.Fprintln(&b, "Bar")

	return &b
}

func main() {
	done := make(chan bool, 1)

	cli := srgnt.CreateProgram("demo")

	foo := cli.AddCommand("foo", Foo, "Prints Foo")
	foo.AddStringFlag("name", "", "Set a name")

	cli.AddCommand("bar", Bar, "Prints Bar")

	cli.Run(done)

	<-done
}
