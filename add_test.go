package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddFiles(t *testing.T) {
	r := require.New(t)
	t.Cleanup(clean_index)

	files := []string{"support/file_for_test.txt", "support/fetched_file.txt"}

	err := addFiles(files)
	if err != nil {
		r.NoError(err)
	}
}

func clean_index() {
	path := filepath.Join(REPO_DIR, INDEX_FILE)
	_, err := os.Stat(path)

	if os.IsExist(err) {
		err := os.Remove(path)
		if err != nil {
			log.Fatal("Error remove index file")
		}
	}
}
