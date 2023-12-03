package main

import (
	"archive/zip"
	"crypto/sha256"
	// "fmt"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func createInitDirs() {
	dirs := [...]string{".git2", ".git2/objects"}

	for _, f := range dirs {
		_, err := os.Stat(f)
		if os.IsNotExist(err) {
			fErr := os.Mkdir(f, 0755)

			if fErr != nil {
				log.Fatal(fErr)
			}
		} else {
			continue
		}
	}
}

func createObjDir(f string) string {
	var defDir = ".git2/objects/"
  // тут у меня не работает

	_, err := os.Stat(f)
	// if os.IsNotExist(err) {
	//   fErr := os.Mkdir(defDir + f, 0755)
	//
	//   if fErr != nil {
	//     log.Fatal(fErr)
	//   }
	// }

	if os.IsExist(err) {
		return (defDir + f)
	}

	return (defDir + f)
}

func initialize() {
	createInitDirs()
}

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

	fDir := createObjDir(sh[0:2])

	zipArc, _ := os.Create(fDir + "/" + sh)
	defer zipArc.Close()

	writer := zip.NewWriter(zipArc)

	w, _ := writer.Create(sh)
	io.Copy(w, f)

	writer.Close()
}

func main() {
	var command = os.Args[1]

	switch command {
	case "init":
		initialize()
	case "-pack":
		if len(os.Args) < 2 {
			print("Argument error")
		} else {
			pack(os.Args[2])
		}
	}
}
