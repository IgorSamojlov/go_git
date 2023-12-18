package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/ini.v1"
)

type commitMessage struct {
	kind               string
	parent             string
	author_name        string
	commiter_name      string
	authoCommitTime    int
	commiterCommitTime int
	message            string
}

func (c *commitMessage) Marshal() {
	// res := fmt.Sprintln()
}

func (c *commitMessage) Unmarshal(message []byte) {
	c.message = string(message)
	c.kind = "object"
	c.parent, _ = getCurrentSha()
	c.author_name, _ = getAuthorName()
	c.commiter_name = c.author_name
}

func getCurrentSha() (string, error) {
	headFile, err := os.Open(filepath.Join(REPO_DIR, HEAD_FILE))
	if err != nil {
		return "", err
	}
	defer headFile.Close()

	fileScanner := bufio.NewScanner(headFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	ref := fileScanner.Text()

	r, err := regexp.Compile("\\w*$")
	if err != nil {
		return "", err
	}

	bFileName := r.FindString(ref)
	branchFile, err := os.Open(filepath.Join(REPO_DIR, REFS_DIR, HEAD_DIR, bFileName))

	fileScanner = bufio.NewScanner(branchFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	sha := fileScanner.Text()

	return sha, nil
}

func getAuthorName() (string, error) {
	cfgDir := filepath.Join(REPO_DIR, CONFIG_FILE)
	cfg, err := ini.Load(cfgDir)
	if err != nil {
		return "", err
	}

	section := cfg.Section("user")
	name := section.Key("name").String()

	return name, nil
}
