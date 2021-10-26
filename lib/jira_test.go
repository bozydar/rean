package lib

import (
	"github.com/joho/godotenv"
	"testing"
)

func Test_GetIssueById(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fail()
	}
	issue := GetIssueById("BT-7021")
	if issue.err != nil {
		t.Fail()
	}
	println(issue.Status)
	if issue.Summary == "" {
		t.Error("Summary is blank")
	}

}
