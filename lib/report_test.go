package lib

import (
	"testing"
)

func Test_extractIssueIds(t *testing.T) {
	result := extractIssueIds("[7021] abc [7022]")
	if result[0] != "7021" {
		t.Error("Bad result")
	}
	if result[1] != "7022" {
		t.Error("Bad result")
	}
}
