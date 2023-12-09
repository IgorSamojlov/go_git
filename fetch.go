package main

import (
	"compress/flate"
	"io"
	"os"
)

type customReadCloser struct {
	flatReader io.ReadCloser
	file       io.ReadCloser
}

func (c customReadCloser) Read(p []byte) (int, error) {
	return c.flatReader.Read(p)
}

func (c customReadCloser) Close() error {
	c.flatReader.Close()
	return c.file.Close()
}

func fetchFile(sum string, filename string) error {
	r, err := fetch(sum)
	if err != nil {
		return err
	}
	defer r.Close()

	file, err := os.Create(fullPath(sum))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	if err != nil {
		return err
	}

	return nil
}

func fetchData(sum string) ([]byte, error) {
	r, err := fetch(sum)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}

func fetch(sum string) (io.ReadCloser, error) {
	storedFile, err := os.Open(fullPath(sum))
	if err != nil {
		return nil, err
	}

	flatReader := flate.NewReader(storedFile)

	return customReadCloser{flatReader: flatReader, file: storedFile}, nil
}
