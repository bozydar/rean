package lib

import (
	"github.com/andygrunwald/go-jira"
	"os"
)

type Issue struct {
	Id      string
	Summary string
	Status  string
	Err     error
}

func GetIssueByIdChannel(id string) (ch chan Issue) {
	ch = make(chan Issue)
	go func() {
		ch <- GetIssueById(id)
	}()
	return
}

func GetIssueById(id string) Issue {
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USERNAME"),
		Password: os.Getenv("JIRA_PASSWORD"),
	}
	jiraClient, _ := jira.NewClient(tp.Client(), "https://1centre.atlassian.net")
	issue, _, err := jiraClient.Issue.Get(id, nil)
	if err != nil {
		return Issue{
			Id:  id,
			Err: err,
		}
	}
	return Issue{
		Id:      id,
		Summary: issue.Fields.Summary,
		Status:  issue.Fields.Status.Name,
		Err:     nil,
	}
}
