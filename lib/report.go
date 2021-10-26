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
	occur := map[string]*Issue{}
	s := spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()
	for _, commit := range commits {
		for _, issueId := range extractIssueIds(commit.Subject) {
			if occur[issueId] == nil {
				issue, err := GetIssueById("BT-" + issueId)
				if err != nil {
					_, err := fmt.Fprintf(os.Stderr, "Can't find %s issue. Reason: %s", issueId, err.Error())
					if err != nil {
						return nil, err
					}
					continue
				}
				occur[issueId] = issue
			}
		}
	}
	s.Stop()

	var result []Issue
	for _, issue := range occur {
		result = append(result, *issue)
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
