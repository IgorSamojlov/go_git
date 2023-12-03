package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	space            = ' '
	openSym          = '<'
	closeSym         = '>'
	plus             = '+'
	minus            = '-'
	parent           = "parent"
	author           = "author"
	commiter         = "committer"
	COMMIT_DATA_SIZE = 512
)

type commitData struct {
	parent             string
	authorName         string
	authorEmail        string
	commiterName       string
	commiterEmail      string
	authorCommitTime   time.Time
	commiterCommitTime time.Time
	message            string
}

func (c *commitData) Unmarshal(commitData []byte) error {
	b := bytes.NewBuffer(commitData)

	if c.parent != "" {
		c.parent = ""
	}

	for {
		line, err := b.ReadString('\n')
		if err != nil {
			break
		}
		switch {
		case strings.HasPrefix(line, author):
			c.authorName, c.authorEmail, c.authorCommitTime, err = parseHuman(line)

			if err != nil {
				return err
			}
		case strings.HasPrefix(line, commiter):
			c.commiterName, c.commiterEmail, c.commiterCommitTime, err = parseHuman(line)
			if err != nil {
				return err
			}
		case strings.HasPrefix(line, parent):
			c.parent = line[7:47]
		case line == "\n":
			c.message = b.String()

			return nil
		}
	}

	return nil
}

func (c *commitData) Marshal() ([]byte, error) {
	b := &bytes.Buffer{}
	b.Grow(COMMIT_DATA_SIZE)

	if c.parent != "" {
		b.WriteString(fmt.Sprintf("%s %s\n", parent, c.parent))
	}

	userLine(b, "author", c.authorName, c.authorEmail, c.authorCommitTime)
	userLine(b, "commiter", c.commiterName, c.commiterEmail, c.commiterCommitTime)

	b.WriteString("\n")
	b.WriteString(c.message)

	return b.Bytes(), nil
}

func userLine(b *bytes.Buffer, prefix, name, email string, t time.Time) {
	b.WriteString(fmt.Sprintf("%s %s <%s> ", prefix, name, email))
	b.WriteString(fmt.Sprintf("%d %s\n", t.Unix(), t.Format("-0700")))
}

func parseHuman(line string) (string, string, time.Time, error) {
	sName := 0
	sEmail := 0
	sTime := 0
	sTimeLoc := 0

	for i, v := range line {
		if v == space && sName == 0 {
			sName = i
		} else if v == openSym {
			sEmail = i
		} else if v == closeSym {
			sTime = i
		} else if v == space && sTime != 0 {
			sTimeLoc = i
		}
	}

	uN := line[(sName + 1):(sEmail - 1)]
	uE := line[(sEmail) : sTime+1]
	sT := line[(sTime + 2):sTimeLoc]
	sTL := line[(sTimeLoc + 1) : len(line)-1]
	tT, err := parseTime(sT, sTL)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return uN, uE, tT, nil
}

func parseTime(sTime, sTimeLoc string) (time.Time, error) {
	uTime64, err := strconv.ParseInt(sTime, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	t := time.Unix(uTime64, 0)
	lT, err := time.Parse("-0700", sTimeLoc)
	if err != nil {
		return time.Time{}, err
	}
	loc := lT.Location()
	return t.In(loc), nil
}
