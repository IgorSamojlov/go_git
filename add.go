package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func add(files []string) error {
	if len(files) == 0 {
		fmt.Println("Select the files to add or > git add .")
	}

	err := addFiles(files)
	if err != nil {
		return err
	}

	return nil
}

func addFiles(files []string) error {
	iP, err := NewIgnorePatterns("support/.git2ignore")
	if err != nil {
		return err
	}

	iD := indexData{IgrnorePatterns: iP}

	_, err = os.Stat(filepath.Join(REPO_DIR, INDEX_FILE))
	if os.IsExist(err) {
		err := iD.fillIndexData()
		if err != nil {
			return err
		}
	}

	err = iD.addEntries(files, 1)
	if err != nil {
		return err
	}

	iD.sortEntries()
	err = iD.dataToIndexFile()

	return nil
}
