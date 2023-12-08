package main

import (
	"compress/flate"
	"io"
	"log"
	"os"
)

const (
	REPO_DIR = ".git2"
	OBJ_DIR  = "objects"
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
	return nil
}

func fetch(sum string, fName string) {
	storedFile, err := os.Open(fullPath(sum))
	defer storedFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	fetchedFile, err := os.Create(fName)
	defer fetchedFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	flatReader := flate.NewReader(storedFile)
	defer flatReader.Close()

	io.Copy(fetchedFile, flatReader)
}

func main() {
	var err error
	var command = os.Args[1]

	switch command {
	case "init":
		err = repoInit()
		if err != nil {
			log.Fatalf("can not create repo: %s", err)
		}
	case "-store":
		if len(os.Args) < 2 {
			print("Argument error")
		} else {
			storeFile(os.Args[2])
		}
	case "-fetch":
		fetch("362abfcf5ed4e6691c278dea8ec4d67f8d9dd8e0a09e674d0e88928b719d4794", "./fetched_file.txt")
	}
}
