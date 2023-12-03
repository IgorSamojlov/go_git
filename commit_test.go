package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseHuman(t *testing.T) {
	r := require.New(t)
	commitString := "author i.samoylov <i.samoylov@NTB37254-CO.local> 1703289600 +0300\n"

	eName, eMail, eTime, err := parseHuman(commitString)
	r.NoError(err)

	eXTime, err := time.Parse(time.RFC3339, "2023-12-23T00:00:00+00:00")
	r.NoError(err)

	location, err := time.LoadLocation("Europe/Moscow")
	r.NoError(err)

	eXTimeS := eXTime.In(location).Format("2006-01-02 15:04:05")
	nETime := eTime.Format("2006-01-02 15:04:05")

	r.Equal("i.samoylov", eName)
	r.Equal("<i.samoylov@NTB37254-CO.local>", eMail)
	r.Equal(eXTimeS, nETime)
}

func TestCommitDataUnmarshal(t *testing.T) {
	r := require.New(t)

	s := "parent 7d95cb9e00cdca9a6a90538868bce9b7d88f8564\n" +
		"author i.samoylov <i.samoylov@ntb37254-co.local> 1701619058 +0300\n" +
		"committer i.samoylov <i.samoylov@ntb37254-co.local> 1701619058 +0300\n" +
		"\n" +
		"function pack"

	cD := commitData{}
	err := cD.Unmarshal([]byte(s))
	r.NoError(err)

	r.Equal(cD.parent, "7d95cb9e00cdca9a6a90538868bce9b7d88f8564")
	r.Equal("i.samoylov", cD.authorName)
	r.Equal("<i.samoylov@ntb37254-co.local>", cD.authorEmail)
	r.Equal("i.samoylov", cD.commiterName)
	r.Equal("<i.samoylov@ntb37254-co.local>", cD.commiterEmail)
	r.Equal("function pack", cD.message)
	r.NotEmpty(cD.authorCommitTime)
	r.NotEmpty(cD.commiterCommitTime)
}

func TestCommitMarshal(t *testing.T) {
	r := require.New(t)

	tT, err := time.Parse(time.RFC3339, "2023-12-23T00:00:00+03:00")
	r.NoError(err)

	cD := commitData{
		parent:             "7d95cb9e00cdca9a6a90538868bce9b7d88f8564",
		authorName:         "i.samoylov",
		authorEmail:        "i.samoylov@ntb37254-co.local",
		commiterName:       "i.samoylov",
		commiterEmail:      "i.samoylov@ntb37254-co.local",
		authorCommitTime:   tT,
		commiterCommitTime: tT,
		message:            "function pack",
	}

	expected := []byte(`parent 7d95cb9e00cdca9a6a90538868bce9b7d88f8564
author i.samoylov <i.samoylov@ntb37254-co.local> 1703278800 +0300
commiter i.samoylov <i.samoylov@ntb37254-co.local> 1703278800 +0300

function pack`)

	v, err := cD.Marshal()

	r.NoError(err)

	r.Equal(string(expected), string(v))
}

func TestUserLine(t *testing.T) {
	r := require.New(t)

	tT, err := time.Parse(time.RFC3339, "2023-12-23T00:00:00+03:00")
	r.NoError(err)

	b := &bytes.Buffer{}

	userLine(b, "author", "i.samoylov", "i.samoylov@ntb37254-co.local", tT)
	r.NoError(err)

	r.Equal(
		"author i.samoylov <i.samoylov@ntb37254-co.local> 1703278800 +0300\n",
		b.String())
}
