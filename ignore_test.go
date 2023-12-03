package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsMatched(t *testing.T) {
	r := require.New(t)

	iP, err := NewIgnorePatterns("support/.gitignore")
	if err != nil {
		r.NoError(err)
	}

	b, err := iP.isMatched(".git1")
	if err != nil {
		r.NoError(err)
	}

	r.True(b, "Is true")

	b, err = iP.isMatched(".git2")
	if err != nil {
		r.NoError(err)
	}

	r.True(b)

	b, err = iP.isMatched("support/hello.txt")
	if err != nil {
		r.NoError(err)
	}

	r.True(b)

	b, err = iP.isMatched("support/foo.txt")
	if err != nil {
		r.NoError(err)
	}

	r.True(b)
}
