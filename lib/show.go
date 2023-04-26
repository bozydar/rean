package lib

import (
	"fmt"
	"os"
	"strings"
)

type ShowArgs struct {
	From         string `arg:"-f,required" help:"Starting sha"`
	To           string `arg:"-t,required" help:"Ending sha"`
	ShowCommits  bool   `arg:"-c,--show-commits" help:"Show all the commits for each reportItem number"`
	ShowAssignee bool   `arg:"-a,--show-assignee" help:"Show assignee for a ticket"`
}

func ShowMain(args *ShowArgs) {
	config, err := Load("./.rean")
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	reportConfig := ReportConfig{
		JiraConfig: &JiraConfig{
			Username:      config.Jira.Username,
			Password:      config.Jira.Password,
			Url:           config.Jira.Url,
			ProjectPrefix: config.Jira.ProjectPrefix,
		},
		Repo: &Repo{Dir: "./"},
		From: args.From,
		To:   args.To,
	}

	reportItems, err := reportConfig.BuildReport()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	for _, reportItem := range reportItems {
		if reportItem.Error() != nil {
			fmt.Printf("%s: ERROR: %s\n", reportItem.Issue.Id, reportItem.Error())
		} else {
			if args.ShowAssignee {
				fmt.Printf("%s: %s: %s: %s\n", reportItem.Issue.Id, reportItem.Issue.Assignee, fmtStatus(reportItem.Issue.Status), reportItem.Issue.Summary)
			} else {
				fmt.Printf("%s: %s: %s\n", reportItem.Issue.Id, fmtStatus(reportItem.Issue.Status), reportItem.Issue.Summary)
			}
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
	if status == "RELEASED" {
		return "\033[32m" + status + "\033[0m"
	} else if status == "PROD READY" || status == "DONE" {
		return "\033[34m" + status + "\033[0m"
	} else {
		return "\033[31m" + status + "\033[0m"
	}
}
