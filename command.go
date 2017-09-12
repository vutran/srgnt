package srgnt

import (
	"flag"
	"io"
)

type CommandFunction func(flags *flag.FlagSet) io.Reader

type Command struct {
	Description string
	Callback    CommandFunction
}
