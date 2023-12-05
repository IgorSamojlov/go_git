package main

import (
	"compress/flate"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
  "bufio"
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

func fullPath(sum string) string {
	return filepath.Join(REPO_DIR, OBJ_DIR, sum[0:2], sum[2:len(sum)])
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

func pack(fileName string) {
	h := sha256.New()
	f, err := os.Open(fileName)

	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatal(err)
	}

	sh := hex.EncodeToString(h.Sum(nil))

	err = mkdir(REPO_DIR, OBJ_DIR, sh[0:2])

	if err != nil {
		log.Fatal(err)
	}

  buf := bufio.NewReader(f)

	store(buf, fullPath(sh))
}

func store(b *bufio.Reader, oFileName string) {
	compressed, err := os.Create(oFileName)
	defer compressed.Close()

	fWriter, err := flate.NewWriter(compressed, flate.NoCompression)

	if err != nil {
		log.Fatal(err)
	}

	defer fWriter.Close()
	io.Copy(fWriter, b)

	fWriter.Flush()
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

func main() {
	var err error
	var command = os.Args[1]

	switch command {
	case "init":
		err = repoInit()
		if err != nil {
			log.Fatalf("can not create repo: %", err)
		}
	case "-store":
		if len(os.Args) < 2 {
			print("Argument error")
		} else {
			pack(os.Args[2])
		}
	case "-fetch":
		fetch("362abfcf5ed4e6691c278dea8ec4d67f8d9dd8e0a09e674d0e88928b719d4794", "./fetched_file.txt")
	}
}
