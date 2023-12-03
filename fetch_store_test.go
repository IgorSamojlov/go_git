package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoreFile(t *testing.T) {
	r := require.New(t)

	t.Cleanup(clean)

	err := repoInit()
	if err != nil {
		t.Errorf("Dir is not created %s", err)
	}

	sha, err := storeFile("support/file_for_test.txt")
	if err != nil {
		t.Errorf("Store file error %s", err)
	}

	fs, err := os.Stat(
		".git2/objects/00/7d3b529e897d9330e542fe5d15ade86fdd1ddf",
	)
	r.NoError(err)

	r.Equal(int64(5308), fs.Size())
	r.NotNil(sha)

	t.Cleanup(clean)
}

func TestStoreFetchFile(t *testing.T) {
	r := require.New(t)

	t.Cleanup(clean)

	err := repoInit()
	if err != nil {
		t.Errorf("Dir is not created %s", err)
	}

	_, err = storeFile("support/file_for_test.txt")
	if err != nil {
		t.Errorf("Store file error %s", err)
	}

	err = fetchFile("007d3b529e897d9330e542fe5d15ade86fdd1ddf", "support/fetched_file.txt")

	r.NoError(err)
}

func TestStoreData(t *testing.T) {
	r := require.New(t)

	t.Cleanup(clean)

	err := repoInit()
	if err != nil {
		t.Errorf("Dir is not created %s", err)
	}

	data := []byte("Data to store")

	sha, err := storeData(data)
	if err != nil {
		t.Errorf("Store file error %s", err)
	}

	fs, err := os.Stat(
		".git2/objects/3c/6cb9cbda027f6ac1f04c9e1b3c756c7cad5ff6",
	)
	r.NoError(err)

	r.Equal(int64(19), fs.Size())
	r.NotNil(sha)

	t.Cleanup(clean)
}
