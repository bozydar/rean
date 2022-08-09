package lib

import (
	"testing"
)

func Test_extractIssueIds(t *testing.T) {
	result := extractIssueIds("[7021] abc [BT-7022123]", "BT")
	if result[0] != "7021" {
		t.Error("Bad result")
	}
	if result[1] != "BT-7022123" {
		t.Error("Bad result", result[1])
	}
}
