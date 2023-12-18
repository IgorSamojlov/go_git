package main

import (
	"os"
)

func clean() {
	os.RemoveAll(".git2")
}

// func TestRepoInit(t *testing.T) {
// 	t.Cleanup(clean)
//
// 	err := repoInit()
// 	if err != nil {
// 		t.Errorf("Dir is not created %s", err)
// 	}
//
// 	dirs := []string{".git2", ".git2/objects", ".git2/refs", ".git2/refs/heads"}
//
// 	for _, s := range dirs {
// 		_, err = os.Stat(s)
//
// 		assert.Equal(t, err, nil, "Dir: %s", s)
// 	}
//
// 	t.Cleanup(clean)
// }
//
// func TestStoreFile(t *testing.T) {
// 	assert := assert.New(t)
//
// 	t.Cleanup(clean)
//
// 	err := repoInit()
// 	err = repoInit()
// 	if err != nil {
// 		t.Errorf("Dir is not created %s", err)
// 	}
//
// 	err = storeFile("file_for_test.txt")
//
// 	if err != nil {
// 		t.Errorf("Store file error %s", err)
// 	}
//
// 	fs, err := os.Stat(
// 		".git2/objects/00/7d3b529e897d9330e542fe5d15ade86fdd1ddf",
// 	)
//
// 	assert.Equal(int64(5308), fs.Size())
//
// 	t.Cleanup(clean)
// }
//
// func TestStoreData(t *testing.T) {
// 	assert := assert.New(t)
//
// 	// t.Cleanup(clean)  непонятно почему но так нельзя
//
// 	err := repoInit()
// 	err = repoInit()
// 	if err != nil {
// 		t.Errorf("Dir is not created %s", err)
// 	}
//
// 	data := []byte("Data to store")
//
// 	err = storeData(data)
// 	if err != nil {
// 		t.Errorf("Store file error %s", err)
// 	}
//
// 	fs, err := os.Stat(
// 		".git2/objects/3c/6cb9cbda027f6ac1f04c9e1b3c756c7cad5ff6",
// 	)
//
// 	assert.Equal(int64(19), fs.Size())
//
// 	t.Cleanup(clean)
// }
