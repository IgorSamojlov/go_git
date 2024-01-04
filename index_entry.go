package main

import (
	"crypto/sha1"
	"encoding/binary"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type indexEntry struct {
	Sha         []byte
	FullName    string
	Permissions uint32
	FileSize    int64
	Status      int
	UpTime      int64
}

func (iE *indexEntry) Fill(path string, f fs.DirEntry, status int) error {
	h := sha1.New()

	iE.FullName = filepath.Join(path, f.Name())

	file, err := os.Open(iE.FullName)
	if err != nil {
		return err
	}

	_, err = io.Copy(h, file)
	if err != nil {
		return err
	}

	iE.Sha = h.Sum(nil)

	fInfo, err := f.Info()
	if err != nil {
		return err
	}

	iE.FileSize = fInfo.Size()
	iE.Permissions = uint32(fInfo.Mode())
	iE.UpTime = fInfo.ModTime().Unix()
	iE.Status = status

	return nil
}

func filledEntry(f *os.File) (error, indexEntry) {
	iE := indexEntry{}

	// sha

	fB := make([]byte, 20)
	_, err := f.Read(fB)
	if err != nil {
		return nil, indexEntry{}
	}
	iE.Sha = fB

	// time

	fB = make([]byte, 8)
	_, err = f.Read(fB)
	if err != nil {
		return nil, indexEntry{}
	}

	iE.UpTime = int64(binary.BigEndian.Uint64(fB))

	// permissions

	fB = make([]byte, 4)
	_, err = f.Read(fB)
	if err != nil {
		return nil, indexEntry{}
	}

	iE.Permissions = binary.BigEndian.Uint32(fB)

	// filesize

	fB = make([]byte, 4)
	_, err = f.Read(fB)
	if err != nil {
		return nil, indexEntry{}
	}

	iE.FileSize = int64(binary.BigEndian.Uint32(fB))

	// pathSize

	fB = make([]byte, 4)
	_, err = f.Read(fB)
	if err != nil {
		return nil, indexEntry{}
	}

	pathSize := int64(binary.BigEndian.Uint32(fB))

	// path

	fB = make([]byte, pathSize)
	_, err = f.Read(fB)
	if err != nil {
		return nil, indexEntry{}
	}

	iE.FullName = string(fB)

	// status

	fB = make([]byte, 2)
	_, err = f.Read(fB)
	if err != nil {
		return nil, indexEntry{}
	}

	iE.Status = int(binary.BigEndian.Uint16(fB))

	return nil, iE
}
