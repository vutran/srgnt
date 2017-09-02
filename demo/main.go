package main

import (
	"flag"
	"fmt"
	"github.com/vutran/srgnt"
)

// TODO - return a map of flags instead (key ,value)
func Foo(flags *flag.FlagSet) {
	name := flags.Lookup("name")

	fmt.Printf("Foo: %s\n", name.Value.String())
}

func Bar(flags *flag.FlagSet) {
	fmt.Println("Bar")
}

func main() {
	cli := srgnt.CreateProgram("demo")

	foo := cli.AddCommand("foo", Foo, "Prints Foo")
	foo.AddStringFlag("name", "", "Set a name")

	cli.AddCommand("bar", Bar, "Prints Bar")
	//bar.AddStringFlag("name", "", "Set a name")

	cli.Run()
}
