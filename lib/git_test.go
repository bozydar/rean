package lib

import (
	"testing"
)

func Test_parseLog(t *testing.T) {
	result := parseLog(`84fdac40d|[6622] - Implementing tests
03eaf3970|[6622] - Email Log: Implementing email_activities endpoint
03eaf3087|[BT-12345|BT-1234] - Something else
`)

	if !(result[0].Sha == "84fdac40d" && result[0].Subject == "[6622] - Implementing tests") {
		t.Errorf("Improper values")
	}
	if !(result[1].Sha == "03eaf3970" && result[1].Subject == "[6622] - Email Log: Implementing email_activities endpoint") {
		t.Errorf("Improper values")
	}
	if !(result[2].Sha == "03eaf3087" && result[2].Subject == "[BT-12345|BT-1234] - Something else") {
		t.Errorf("Improper values")
	}
}
