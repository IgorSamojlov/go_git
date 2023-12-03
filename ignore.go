package main

import (
	"bytes"
	"os"
	"path/filepath"
)

type ignorePatterns struct {
	patterns [][]byte
}

func NewIgnorePatterns(filename string) (*ignorePatterns, error) {
	i := &ignorePatterns{}

	fB, err := os.ReadFile(filename)
	if err != nil {
		return i, err
	}
	i.patterns = bytes.Split(fB, []byte("\n"))

	return i, nil
}

func (i ignorePatterns) isMatched(filename string) (bool, error) {
	for _, p := range i.patterns {
		matched, err := filepath.Match(string(p), filename)
		if err != nil {
			return false, err
		}

		if matched {
			return true, nil
		}
	}
	return false, nil
}
