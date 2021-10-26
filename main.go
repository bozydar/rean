package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"os"
	"rean/lib"
)

func main() {
	var args struct {
		Dir  string `arg:"positional" placeholder:"WOR_DIR" help:"Where the analysed directory is" default:"."`
		From string `arg:"--from" help:"Starting sha"`
		To   string `arg:"--to" help:"Ending sha"`
	}
	arg.MustParse(&args)
	repo := lib.Repo{
		Dir: args.Dir,
	}
	issues, err := lib.BuildReport(&repo, args.From, args.To)
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	for _, issue := range issues {
		fmt.Printf("%+v\n", issue)
	}
}
