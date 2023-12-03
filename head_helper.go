package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
)

func getCurrentSha() (string, error) {
	bFileName, err := getCurrentRef()
	if err != nil {
		return "", err
	}

	branchFile, err := os.Open(filepath.Join(REPO_DIR, REFS_DIR, HEAD_DIR, bFileName))
	if err != nil {
		return "", err
	}

	fileScanner := bufio.NewScanner(branchFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	sha := fileScanner.Text()

	return sha, nil
}

func setCurrentSha(sha string) error {
	bFileName, err := getCurrentRef()
	if err != nil {
		return err
	}

	err = os.WriteFile(
		filepath.Join(REPO_DIR, REFS_DIR, HEAD_DIR, bFileName),
		[]byte(sha),
		0644,
	)

	if err != nil {
		return err
	}

	return nil
}

func getCurrentRef() (string, error) {
	headFile, err := os.Open(filepath.Join(REPO_DIR, HEAD_FILE))
	if err != nil {
		return "", err
	}
	defer headFile.Close()

	fileScanner := bufio.NewScanner(headFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	ref := fileScanner.Text()

	r, err := regexp.Compile(`\w*$`)
	if err != nil {
		return "", err
	}

	return r.FindString(ref), nil
}
