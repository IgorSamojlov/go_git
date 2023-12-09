package main

import (
	"os"
	"testing"
)

func clean() {
	os.RemoveAll(".git2")
}

func TestRepoInit(t *testing.T) {
	_ = repoInit()
	_, err := os.Stat(".git2")
	if err != nil {
		t.Errorf("Dir is not created %s", err)
	}
}

func TestStoreFile(t *testing.T) {
	err := storeFile("file_for_test.txt")
	if err != nil {
		t.Errorf("Store file error %s", err)
	}

	fs, err := os.Stat(
		".git2/objects/80/02935a6c532ea6aed1c45b9ca7f0cb0f1416a7f9e5e3c6966b5c79af44c3dd",
	)
	if fs.Size() != 5308 {
		t.Errorf("Stored file size invalid")
	}
}

func TestStoreData(t *testing.T) {
	data := []byte("Data to store")
	err := storeData(data)
	if err != nil {
		t.Errorf("Store file error %s", err)
	}

	fs, err := os.Stat(
		".git2/objects/e6/d07aa5ac89f43e377fa383a039baeb322d6a6c91fb765d6a469d3e079ea057",
	)

	if fs.Size() != 19 {
		t.Errorf("Stored file size invalid")
	}

	t.Cleanup(clean)
}
