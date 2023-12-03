package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func clean() {
	_, err := os.Stat("/path/to/whatever")
	if os.IsExist(err) {
		os.RemoveAll(".git2")
	}
}

func TestRepoInit(t *testing.T) {
	t.Cleanup(clean)

	err := repoInit()
	if err != nil {
		t.Errorf("Dir is not created %s", err)
	}

	dirs := []string{".git2", ".git2/objects", ".git2/refs", ".git2/refs/heads"}

	for _, s := range dirs {
		_, err = os.Stat(s)

		require.Equal(t, err, nil, "Dir: %s", s)
	}

	t.Cleanup(clean)
}

func TestUserData(t *testing.T) {
	r := require.New(t)

	err := repoInit()
	r.NoError(err)

	err = configure("user.name", "test_name")
	r.NoError(err)

	err = configure("user.email", "test@test.ru")
	r.NoError(err)

	user, err := getUserData()
	r.NoError(err)

	r.Equal(user.Name, "test_name")
}
