package main

import (
	"bytes"
	"compress/flate"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func storeFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return store(f)
}

func storeData(data []byte) error {
	r := bytes.NewReader(data)
	return store(r)
}

// store - save byte array into the storage
func store(data io.ReadSeeker) error {
	h := sha1.New()

	_, err := io.Copy(h, data)
	if err != nil {
		return err
	}

	_, err = data.Seek(0, 0)
	if err != nil {
		return err
	}

	sh := hex.EncodeToString(h.Sum(nil))

	err = mkdir(REPO_DIR, OBJ_DIR, sh[0:2])
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath(sh))
	if err != nil {
		return err
	}

	defer file.Close()

	w, err := flate.NewWriter(file, flate.BestCompression)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, data)
	if err != nil {
		return err
	}

	return w.Flush()
}
