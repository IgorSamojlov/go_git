package main

import (
	"bytes"
	"fmt"
	"time"
)

const dateLayout = "Date:   " + time.RFC1123Z

func showLog(c string) error {
	cD := &commitData{}

	sha, err := getCurrentSha()
	if err != nil {
		return err
	}

	for sha != "" {
		data, err := fetchData(sha)
		if err != nil {
			return err
		}

		err = cD.Unmarshal(data)
		if err != nil {
			return err
		}

		drawCommitContent(cD, sha)
		sha = cD.parent
	}

	return nil
}

func drawCommitContent(c *commitData, sha string) {
	b := &bytes.Buffer{}

	b.WriteString(fmt.Sprintf("%scommmit %s %s\n", "\033[33m", sha, "\033[0m"))
	b.WriteString(fmt.Sprintf("Author: %s %s\n", c.authorName, c.authorEmail))
	b.WriteString(c.authorCommitTime.Format(dateLayout))
	b.WriteString("\n")
	b.WriteString("     ")
	b.WriteString(c.message)
	b.WriteString("\n")

	fmt.Println(b.String())
}
