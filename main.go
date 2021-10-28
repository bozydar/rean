package main

import (
	"github.com/alexflint/go-arg"
	"rean/lib"
)

func main() {
	var args lib.Args
	arg.MustParse(&args)

	switch {
	case args.Config != nil:
		lib.ConfigMain(args.Config)
	case args.Show != nil:
		lib.ShowMain(args.Show)
	}
}
