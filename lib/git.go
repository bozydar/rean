package lib

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
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
	cmd := exec.Command("git", "log", fmt.Sprintf("%s..%s", from, to), "--pretty=format:%h|%s")
	cmd.Dir = repo.Dir

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(&stdBuffer)

	cmd.Stdout = mw

	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return parseLog(stdBuffer.String()), nil
}

func parseLog(log string) (result []Commit) {
	lines := strings.Split(log, "\n")
	result = []Commit{}
	for _, line := range lines {
		shaAndSubject := strings.Split(line, "|")
		result = append(result, Commit{
			Sha:     shaAndSubject[0],
			Subject: shaAndSubject[1],
		})
	}
	return result
}
