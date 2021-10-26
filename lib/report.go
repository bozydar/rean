package lib

import (
	"fmt"
	"github.com/briandowns/spinner"
	"os"
	"regexp"
	"time"
)

func BuildReport(repo *Repo, from string, to string) ([]Issue, error) {
	commits, err := repo.Diff(from, to)
	if err != nil {
		return nil, err
	}
	channelById := map[string]chan Issue{}
	s := spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()
	for _, commit := range commits {
		for _, issueId := range extractIssueIds(commit.Subject) {
			if channelById[issueId] == nil {
				channelById[issueId] = GetIssueByIdChannel("BT-" + issueId)
			}
		}
	}
	s.Stop()
	var result []Issue
	for issueId, requestChannel := range channelById {
		issue, _ := <-requestChannel

		if issue.err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can't find %s issue. Reason: %s", issueId, issue.err.Error())
			continue
		}
		result = append(result, Issue(issue))
	}

	return result, nil
}

var issueIdRe = regexp.MustCompile(`\[(\d+?)]`)

func extractIssueIds(subject string) []string {
	found := issueIdRe.FindAllString(subject, -1)
	if found == nil {
		return []string{}
	}
	var result []string
	for _, value := range found {
		result = append(result, value[1:len(value)-1])
	}
	return result
}
