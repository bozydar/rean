package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"os"
	"rean/lib"
	"strings"
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
	println()
	for _, issue := range issues {
		if issue.Err != nil {
			fmt.Printf("%s: ERROR: %s\n", issue.Id, issue.Err.Error())
		} else {
			fmt.Printf("%s: %s: %s\n", issue.Id, fmtStatus(issue.Status), issue.Summary)
		}
	}
}

func fmtStatus(status string) string {
	status = strings.ToUpper(status)
	if status == "PROD READY" || status == "RELEASED" {
		return "\033[32m" + status + "\033[0m"
	} else {
		return "\033[31m" + status + "\033[0m"
	}
}
