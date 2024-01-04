package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type treeEntry struct {
	Tree bool
	Sha  string
	Name string
	Mode fs.FileMode
}

type treeData []treeEntry

func (d treeData) Marshal() ([]byte, error) {
	b := &bytes.Buffer{}
	b.Grow(len(d) * 100)

	t := ""

	for _, entry := range d {
		if entry.Tree {
			t = "tree"
		} else {
			t = "blob"
		}
		b.WriteString(fmt.Sprintf("%6o %s %s\t%s\n", entry.Mode, t, entry.Sha, entry.Name))
	}

	return b.Bytes(), nil
}

func tree(path string, iP *ignorePatterns) (string, error) {
	isM, err := iP.isMatched(path)
	if err != nil {
		return "", err
	}
	if isM {
		fmt.Println(path)
		return "", nil
	}

	b := treeData{}
	if path == "" {
		path = "./"
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		isM, err := iP.isMatched(file.Name())
		if err != nil {
			return "", err
		}
		if isM {
			continue
		}

		sha := ""
		info, err := file.Info()
		if err != nil {
			return "", err
		}

		if file.IsDir() {
			sha, err = tree(filepath.Join(path, file.Name()), iP)
		} else {
			sha, err = storeFile(filepath.Join(path, file.Name()))
		}
		if err != nil {
			return "", err
		}

		b = append(b, treeEntry{Tree: file.IsDir(), Sha: sha, Name: file.Name(), Mode: info.Mode()})
	}

	tData, err := b.Marshal()
	if err != nil {
		return "", err
	}

	sha, err := storeData(tData)
	if err != nil {
		return "", err
	}

	return sha, nil
}
