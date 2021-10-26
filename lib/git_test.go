package lib

import (
	"testing"
)

func Test_parseLog(t *testing.T) {
	result := parseLog(`84fdac40d|[6622] - Implementing tests
03eaf3970|[6622] - Email Log: Implementing email_activities endpoint`)

	if !(result[0].Sha == "84fdac40d" && result[0].Subject == "[6622] - Implementing tests") {
		t.Errorf("Improper values")
	}
	if !(result[1].Sha == "03eaf3970" && result[1].Subject == "[6622] - Email Log: Implementing email_activities endpoint") {
		t.Errorf("Improper values")
	}
}
