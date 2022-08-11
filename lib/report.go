package lib

import (
	"fmt"
	"regexp"
	"strings"
)

type ReportItem struct {
	Issue   Issue
	Commits []Commit
}

type commitAndCh struct {
	ch      chan Issue
	commits []Commit
}

type ReportConfig struct {
	JiraConfig *JiraConfig
	Repo       *Repo
	From       string
	To         string
}

func (reportItem *ReportItem) Error() error {
	return reportItem.Issue.Err
}

func (reportConfig *ReportConfig) BuildReport() ([]ReportItem, error) {
	commits, err := reportConfig.Repo.Diff(reportConfig.From, reportConfig.To)
	if err != nil {
		return nil, err
	}
	channelById := map[string]*commitAndCh{}
	for _, commit := range commits {
		for _, issueId := range extractIssueIds(commit.Subject, reportConfig.JiraConfig.ProjectPrefix) {
			if channelById[issueId] == nil {
				channelById[issueId] = &commitAndCh{
					ch:      reportConfig.JiraConfig.GetIssueByIdChannel(issueId),
					commits: []Commit{commit},
				}
			} else {
				channelById[issueId].commits = append(channelById[issueId].commits, commit)
			}
		}
	}
	var result []ReportItem
	for _, commitAndCh := range channelById {
		issue, _ := <-commitAndCh.ch
		reportItem := ReportItem{
			Issue:   issue,
			Commits: commitAndCh.commits,
		}
		result = append(result, reportItem)
	}

	return result, nil
}

func extractIssueIds(subject string, projectPrefix string) []string {
	re := fmt.Sprintf(`\[(%s-)?(\d+?)]`, projectPrefix)
	var issueIdRe = regexp.MustCompile(re)

	found := issueIdRe.FindAllString(subject, -1)
	if found == nil {
		return []string{}
	}
	var result []string
	for _, value := range found {
		id := buildIdFromNumber(value[1:len(value)-1], projectPrefix)
		result = append(result, id)
	}
	return result
}

func buildIdFromNumber(number string, projectPrefix string) string {
	if strings.Contains(number, "-") {
		return number
	} else {
		return projectPrefix + "-" + number
	}
}
