package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var err error
	command := os.Args[1]
	iChecker, err := NewIgnorePatterns(".git2ignore")
	if err != nil {
		log.Fatal("Can nod load ignore checher")
	}

	switch command {
	case "init":
		err = repoInit()
		if err != nil {
			log.Fatalf("Can not create repo: %s", err)
		}
	case "config":
		err = configure(os.Args[2], os.Args[3])
		if err != nil {
			log.Fatalf("Can not configure: %s", err)
		}
	case "commit":
		err = commit(os.Args[2])
		if err != nil {
			log.Fatalf("Can not commit: %s", err)
		}
	case "log":
		err = showLog(os.Args[2])
		if err != nil {
			log.Fatalf("Can not log: %s", err)
		}
	case "tree":
		sha, err := tree(os.Args[2], iChecker)
		if err != nil {
			log.Fatalf("Can not log: %s", err)
		}
		fmt.Println(sha)
	case "cat-file":
		data, err := fetchData(os.Args[2])
		if err != nil {
			log.Fatalf("Can not log: %s", err)
		}
		fmt.Println(string(data))
	default:
		log.Fatalf("Argument error")
	}
}
