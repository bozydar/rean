package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/briandowns/spinner"
	"os"
	"rean/lib"
	"strings"
	"time"
)

func main() {
	var args struct {
		JiraUsername      string `arg:"--jira-username,env:JIRA_USERNAME"`
		JiraPassword      string `arg:"--jira-password,env:JIRA_PASSWORD"`
		JiraUrl           string `arg:"--jira-url,env:JIRA_URL"`
		JiraProjectPrefix string `arg:"--jira-project-prefix,env:JIRA_PROJECT_PREFIX"`
		Dir               string `arg:"positional" placeholder:"WORK_DIR" help:"Where the analysed directory is" default:"."`
		From              string `arg:"--from" help:"Starting sha"`
		To                string `arg:"--to" help:"Ending sha"`
		ShowCommits       bool   `arg:"-c,--show-commits" help:"Show all the commits for each reportItem number"`
	}
	arg.MustParse(&args)

	s := spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()

	reportConfig := lib.ReportConfig{
		JiraConfig: &lib.JiraConfig{
			Username:      args.JiraUsername,
			Password:      args.JiraPassword,
			Url:           args.JiraUrl,
			ProjectPrefix: args.JiraProjectPrefix,
		},
		Repo: &lib.Repo{Dir: args.Dir},
		From: args.From,
		To:   args.To,
	}

	reportItems, err := reportConfig.BuildReport()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	s.Stop()
	println()
	for _, reportItem := range reportItems {
		if reportItem.Error() != nil {
			fmt.Printf("%s: ERROR: %s\n", reportItem.Issue.Id, reportItem.Error())
		} else {
			fmt.Printf("%s: %s: %s\n", reportItem.Issue.Id, fmtStatus(reportItem.Issue.Status), reportItem.Issue.Summary)
			if args.ShowCommits {
				for _, commit := range reportItem.Commits {
					fmt.Printf("\t%s\n", commit.Subject)
				}
			}
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
