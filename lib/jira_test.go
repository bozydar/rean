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
	if issue.Err != nil {
		t.Fail()
	}
	println(issue.Status)
	if issue.Summary == "" {
		t.Error("Summary is blank")
	}

}
