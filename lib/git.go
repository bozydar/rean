package lib

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Repo struct {
	Dir string
}
type Commit struct {
	Sha     string
	Subject string
}

func (repo *Repo) Diff(from string, to string) ([]Commit, error) {
	args := []string{
		"log",
		fmt.Sprintf("%s..%s", from, to),
		"--pretty=format:%h|%s",
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = repo.Dir

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(&stdBuffer)

	cmd.Stdout = mw

	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Problem with running git "+strings.Join(args, " "))
		return nil, err
	}
	return parseLog(stdBuffer.String()), nil
}

var shaRegexp = regexp.MustCompile(`^([0-9a-f]{9})\|(.*)$`)

func parseLog(log string) (result []Commit) {
	lines := strings.Split(log, "\n")
	result = []Commit{}
	for _, line := range lines {
		groups := shaRegexp.FindAllStringSubmatch(line, -1)
		for _, group := range groups {
			result = append(result, Commit{
				Sha:     group[1],
				Subject: group[2],
			})
		}
	}
	return result
}
