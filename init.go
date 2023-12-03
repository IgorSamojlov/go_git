package main

import (
	"os"
	"path/filepath"
)

const (
	REPO_DIR    = ".git2"
	OBJ_DIR     = "objects"
	HEAD_DIR    = "heads"
	CONFIG_FILE = "config"
	REFS_DIR    = "refs"
	REFS_FILE   = "refs"
	MAIN_BRANCH = "main"
	HEAD_FILE   = "HEAD"
)

func repoInit() error {
	err := mkdir(REPO_DIR)
	if err != nil {
		return err
	}
	err = mkdir(REPO_DIR, OBJ_DIR)
	if err != nil {
		return err
	}
	err = mkdir(REPO_DIR, REFS_DIR)
	if err != nil {
		return err
	}
	err = mkdir(REPO_DIR, REFS_DIR, HEAD_DIR)
	if err != nil {
		return err
	}

	err = createBrancFile()
	if err != nil {
		return err
	}
	err = createHeadFile()
	if err != nil {
		return err
	}

	err = createConfigFile()
	if err != nil {
		return err
	}

	err = createIgnoreFile()
	if err != nil {
		return err
	}

	return nil
}

func createBrancFile() error {
	f, err := os.Create(filepath.Join(REPO_DIR, REFS_DIR, HEAD_DIR, MAIN_BRANCH))
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

func createHeadFile() error {
	f, err := os.Create(filepath.Join(REPO_DIR, HEAD_FILE))
	if err != nil {
		return err
	}
	defer f.Close()

	// for example
	_, err = f.WriteString("ref: refs/head/main")
	if err != nil {
		return err
	}

	return nil
}

func createConfigFile() error {
	f, err := os.Create(filepath.Join(REPO_DIR, CONFIG_FILE))
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

func createIgnoreFile() error {
	f, err := os.Create(".git2ignore")
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}
