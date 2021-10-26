package lib

import (
	jira "github.com/andygrunwald/go-jira"
	"os"
)

type Issue struct {
	Id      string
	Summary string
	Status  string
}

func GetIssueById(id string) (*Issue, error) {
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USERNAME"),
		Password: os.Getenv("JIRA_PASSWORD"),
	}
	jiraClient, _ := jira.NewClient(tp.Client(), "https://1centre.atlassian.net")
	issue, _, err := jiraClient.Issue.Get(id, nil)
	if err != nil {
		return nil, err
	}
	return &Issue{
		Id:      id,
		Summary: issue.Fields.Summary,
		Status:  issue.Fields.Status.Name,
	}, nil
}
