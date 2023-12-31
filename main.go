package main

import (
	"log"
	"os"
)

func main() {
	var err error
	command := os.Args[1]

	switch command {
	case "init":
		err = repoInit()
		if err != nil {
			log.Fatalf("can not create repo: %s", err)
		}
	case "-store":
		if len(os.Args) < 2 {
			print("Argument error")
		} else {
			storeFile(os.Args[2])
		}
	case "-fetch":
		print("fetch")
	case "config":
		configure(os.Args[2], os.Args[3])
	}
}
