package main

import (
	"archive/zip"
	"crypto/sha256"
	"fmt"
	"path/filepath"

	// "fmt"
	"encoding/hex"
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

func mkdir(names ...string) error {
	name := filepath.Join(names...)

	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return os.Mkdir(name, 0755)
	}

	if info.IsDir() {
		return nil
	}

	return fmt.Errorf("%s is no a directory", name)
}

// fetch()

// store
func pack(fileName string) {
	h := sha256.New()
	f, err := os.Open(fileName)

	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	sh := hex.EncodeToString(h.Sum(nil))
	fDir := filepath.Join(REPO_DIR, OBJ_DIR, sh[0:2])

	zipArc, _ := os.Create(filepath.Join(fDir, sh[2:]))
	defer zipArc.Close()

	writer := zip.NewWriter(zipArc)

	// zlib.Deflate
	w, _ := writer.Create(sh)
	io.Copy(w, f)

	writer.Close()
}

func main() {
	var err error
	var command = os.Args[1]

	switch command {
	case "init":
		err = repoInit()
		if err != nil {
			log.Fatalf("can not create repo: %", err)
		}

	case "-pack":
		if len(os.Args) < 2 {
			print("Argument error")
		} else {
			pack(os.Args[2])
		}
	}
}
