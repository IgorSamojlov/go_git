package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

type indexData struct {
	IndexEntries    []indexEntry
	IgrnorePatterns *ignorePatterns
}

func (iD *indexData) Len() int {
	return len(iD.IndexEntries)
}

func (iD *indexData) sortEntries() {
	sort.SliceStable(iD.IndexEntries, func(i, j int) bool {
		return iD.IndexEntries[i].FullName < iD.IndexEntries[j].FullName
	})
}

// fill data struct from dir

func (iD *indexData) readDir(path string) error {
	err := iD.fillData(path, iD.IgrnorePatterns)
	if err != nil {
		return err
	}

	return nil
}

func (iD *indexData) addEntries(files []string, status int) error {
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		i, err := f.Stat()
		if err != nil {
			return err
		}

		fDE := fs.FileInfoToDirEntry(i)
		path := filepath.Dir(file)

		iE := indexEntry{}
		err = iE.Fill(path, fDE, status)
		if err != nil {
			return err
		}

		iD.IndexEntries = append(iD.IndexEntries, iE)
	}

	return nil
}

func (iD *indexData) fillData(path string, iP *ignorePatterns) error {
	dFiles, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range dFiles {
		isM, err := iP.isMatched(f.Name())
		if err != nil {
			return err
		}

		isMD, err := iP.isMatched(filepath.Join(path, f.Name()))
		if err != nil {
			return err
		}

		if isM || isMD {
			continue
		}

		if f.IsDir() {
			err := iD.fillData(filepath.Join(path, f.Name()), iP)
			if err != nil {
				return err
			}

			continue
		}

		entry := indexEntry{}
		err = entry.Fill(path, f, 0)
		if err != nil {
			return err
		}

		iD.IndexEntries = append(iD.IndexEntries, entry)
	}
	return nil
}

// fill data struct from index file

func (iD *indexData) fillIndexData() error {
	file, err := os.Open(filepath.Join(REPO_DIR, INDEX_FILE))
	if err != nil {
		return err
	}

	fB := make([]byte, 2)
	_, err = file.Read(fB)
	if err != nil {
		return err
	}

	size := int(binary.BigEndian.Uint16(fB))

	for size != 0 {
		err, e := filledEntry(file)
		if err != nil {
			return err
		}

		iD.IndexEntries = append(iD.IndexEntries, e)

		size--
	}

	return nil
}

// from data struct to file

func (iD *indexData) dataToIndexFile() error {
	b := &bytes.Buffer{}
	size := make([]byte, 2)
	binary.BigEndian.PutUint16(size, uint16((len(iD.IndexEntries))))
	b.Write(size)

	for _, entry := range iD.IndexEntries {
		err := addSha(entry.FullName, b)
		if err != nil {
			return err
		}

		err = addFileInfo(entry, b)
		if err != nil {
			return err
		}

		err = addFullName(entry.FullName, b)
		if err != nil {
			return err
		}

		err = addState(entry.Status, b)
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(REPO_DIR, INDEX_FILE), b.Bytes(), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// helper functions

func addSha(path string, b *bytes.Buffer) error {
	h := sha1.New()

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(h, f)
	if err != nil {
		return err
	}

	f.Close()

	_, err = b.Write(h.Sum(nil))
	if err != nil {
		return err
	}

	return nil
}

func addFileInfo(iE indexEntry, b *bytes.Buffer) error {
	tB := make([]byte, 8)
	binary.BigEndian.PutUint64(tB, uint64(iE.UpTime))
	_, err := b.Write(tB)
	if err != nil {
		return err
	}

	tB = make([]byte, 4)
	binary.BigEndian.PutUint32(tB, uint32(iE.Permissions))
	_, err = b.Write(tB)

	if err != nil {
		return err
	}

	tB = make([]byte, 4)
	binary.BigEndian.PutUint32(tB, uint32(iE.FileSize))
	_, err = b.Write(tB)
	if err != nil {
		return err
	}

	return nil
}

func addFullName(fullName string, b *bytes.Buffer) error {
	tB := make([]byte, 4)
	binary.BigEndian.PutUint32(tB, uint32(len(fullName)))
	_, err := b.Write(tB)
	if err != nil {
		return err
	}

	_, err = b.WriteString(fullName)
	if err != nil {
		return err
	}

	return nil
}

func addState(s int, b *bytes.Buffer) error {
	tB := make([]byte, 2)
	binary.BigEndian.PutUint16(tB, uint16(s))
	_, err := b.Write(tB)
	if err != nil {
		return err
	}

	return nil
}
