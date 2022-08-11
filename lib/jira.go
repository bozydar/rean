package lib

import (
	"github.com/andygrunwald/go-jira"
)

type Issue struct {
	Id      string
	Summary string
	Status  string
	Err     error
}

type JiraConfig struct {
	Username      string
	Password      string
	Url           string
	ProjectPrefix string
}

func (jiraConfig *JiraConfig) GetIssueByIdChannel(id string) (ch chan Issue) {
	ch = make(chan Issue)
	go func() {
		ch <- jiraConfig.GetIssueById(id)
	}()
	return
}

func (jiraConfig *JiraConfig) GetIssueById(id string) Issue {
	jiraClient, _ := jiraConfig.buildJiraClient()
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

func (jiraConfig *JiraConfig) buildJiraClient() (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: jiraConfig.Username,
		Password: jiraConfig.Password,
	}
	return jira.NewClient(tp.Client(), jiraConfig.Url)
}
