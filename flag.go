package srgnt

import (
	"flag"
	"time"
)

func (p *Program) AddBoolFlag(name string, value bool, usage string) {
	flag.Bool(name, value, usage)
}

func (p *Program) AddDurationFlag(name sting, value time.Duration, usage string) {
	flag.Duration(name, value, usage)
}

func (p *Program) AddFloat64Flag(name sting, value float64, usage string) {
	flag.Float64(name, value, usage)
}

func (p *Program) AddIntFlag(name string, value int, usage string) {
	flag.Int(name, value, usage)
}

func (p *Program) AddInt64Flag(name sting, value int64, usage string) {
	flag.Int64(name, value, usage)
}

func (p *Program) AddStringFlag(name string, value string, usage string) {
	flag.String(name, value, usage)
}

func (p *Program) AddUint(name sting, value uint, usage string) {
	flag.Uint(name, value, usage)
}

func (p *Program) AddUint64(name sting, value uint64, usage string) {
	flag.Uint64(name, value, usage)
}
