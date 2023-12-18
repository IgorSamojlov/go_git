package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitMessageUnmurshal(t *testing.T) {
	assert := assert.New(t)
	prepare(t)

	cM := commitMessage{}
	message := []byte("commit message")
	cM.Unmarshal(message)

	assert.Equal(cM.message, string(message))
	assert.Equal(cM.kind, "object")
	assert.Equal(cM.parent, "b28b7af69320201d1cf206ebf28373980add1451")
	assert.Equal(cM.author_name, "test_name")
	assert.Equal(cM.commiter_name, "test_name")
}

func prepare(t *testing.T) {
	err := repoInit()
	if err != nil {
		t.Errorf("Dir is not created %s", err)
	}

	data := []byte("b28b7af69320201d1cf206ebf28373980add1451")
	err = os.WriteFile(".git2/refs/heads/main", data, 0644)

	if err != nil {
		t.Errorf("Error on prepare %s", err)
	}

	configure("user.name", "test_name")
	configure("user.email", "test@test.ru")
}
