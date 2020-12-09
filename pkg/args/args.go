package args

import (
	"flag"
)

type Args struct {
	File   string
	Part   int
	Number int
}

var (
	args *Args
	part = flag.Int("part", 1, "which part to run")
	num  = flag.Int("n", 0, "a number param")
)

func ParseArgs() Args {
	if args != nil {
		return *args
	}

	flag.Parse()
	args := &Args{
		File:   flag.Arg(0),
		Part:   *part,
		Number: *num,
	}

	return *args
}
