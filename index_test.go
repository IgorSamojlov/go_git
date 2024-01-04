package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexReadDir(t *testing.T) {
	r := require.New(t)

	iP, err := NewIgnorePatterns("support/.git2ignore")
	if err != nil {
		r.NoError(err)
	}

	iD := indexData{IgrnorePatterns: iP}
	err = iD.readDir("support")
	if err != nil {
		r.NoError(err)
	}

	r.Equal(6, len(iD.IndexEntries))
}

func TestWriteToIndexFile(t *testing.T) {
	r := require.New(t)

	iP, err := NewIgnorePatterns("support/.git2ignore")
	if err != nil {
		r.NoError(err)
	}

	iD := indexData{IgrnorePatterns: iP}
	err = iD.readDir("support")
	if err != nil {
		r.NoError(err)
	}

	err = iD.dataToIndexFile()
	if err != nil {
		r.NoError(err)
	}
}

func TestLoadIndexFile(t *testing.T) {
	r := require.New(t)

	iP, err := NewIgnorePatterns("support/.git2ignore")
	if err != nil {
		r.NoError(err)
	}

	iD := indexData{IgrnorePatterns: iP}
	err = iD.readDir("support")
	if err != nil {
		r.NoError(err)
	}

	err = iD.dataToIndexFile()
	if err != nil {
		r.NoError(err)
	}

	iD.IndexEntries = iD.IndexEntries[:0]

	err = iD.fillIndexData()
	if err != nil {
		r.NoError(err)
	}

	r.Equal(6, len(iD.IndexEntries))
	e := iD.IndexEntries[1]
	r.Equal(int64(16), e.FileSize)
	r.Equal(0, e.Status)
}

func TestLoadAddFiles(t *testing.T) {
	r := require.New(t)

	iP, err := NewIgnorePatterns("support/.git2ignore")
	if err != nil {
		r.NoError(err)
	}

	iD := indexData{IgrnorePatterns: iP}
	files := []string{"support/fetched_file.txt", "support/file_for_test.txt"}

	err = iD.addEntries(files, 0)
	if err != nil {
		r.NoError(err)
	}

	r.Equal(2, iD.Len())

	e := iD.IndexEntries[0]
	r.Equal("support/fetched_file.txt", e.FullName)
	r.Equal(0, e.Status)
}
