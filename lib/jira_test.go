package lib

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func Test_GetIssueById(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fail()
	}
	jiraConfig := JiraConfig{
		Username:      os.Getenv("JIRA_USERNAME"),
		Password:      os.Getenv("JIRA_PASSWORD"),
		Url:           os.Getenv("JIRA_URL"),
		ProjectPrefix: os.Getenv("JIRA_PROJECT_PREFIX"),
	}
	issue := jiraConfig.GetIssueById("7021")
	if issue.Err != nil {
		t.Fail()
	}
	println(issue.Status)
	if issue.Summary == "" {
		t.Error("Summary is blank")
	}

}
