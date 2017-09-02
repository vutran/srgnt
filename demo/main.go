package main

import (
	"fmt"
	"github.com/vutran/srgnt"
)

func Foo() {
	fmt.Println("Foo")
}

func Bar() {
	fmt.Println("Bar")
}

func main() {
	cli := srgnt.CreateProgram("demo")
	cli.AddCommand("foo", Foo, "Prints Foo")
	cli.AddCommand("bar", Bar, "Prints Bar")
	cli.Run()
}
