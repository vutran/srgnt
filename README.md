# srgnt

> Experimental CLI Framework

## Install

To get started, you need a working Go environment. Once available, grab the package here:

```
$ go get github.com/vutran/srgnt
```

## Usage

`srgnt` provides a very simple API. The following is an example of a minimal CLI app that has 1 command called "hello" which will print out "Hello, world!".

```
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
```

The following example will output "Hello, world!":

```
$ foo hello
```

## License

MIT © [Vu Tran](https://github.com/vutran/srgnt)
