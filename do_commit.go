package main

import (
	"errors"
	"time"
)

func commit(message string) error {
	if message == "" {
		return errors.New("Empty commit message")
	}

	timeNow := time.Now()

	config, err := readFromConfig()
	if err != nil {
		return err
	}
	sha, err := getCurrentSha()
	if err != nil {
		return err
	}

	cData := commitData{
		parent:             sha,
		authorName:         config.User.Name,
		authorEmail:        config.User.Email,
		authorCommitTime:   timeNow,
		commiterName:       config.User.Name,
		commiterEmail:      config.User.Email,
		commiterCommitTime: timeNow,
		message:            message,
	}

	mData, err := cData.Marshal()
	if err != nil {
		return err
	}

	sha, err = storeData(mData)
	if err != nil {
		return err
	}

	err = setCurrentSha(sha)
	if err != nil {
		return err
	}

	return nil
}
