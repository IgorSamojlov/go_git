package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func fullPath(sum string) string {
	return filepath.Join(REPO_DIR, OBJ_DIR, sum[0:2], sum[2:])
}

func mkdir(names ...string) error {
	name := filepath.Join(names...)

	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return os.Mkdir(name, 0755)
	}

	if info.IsDir() {
		return nil
	}

	return fmt.Errorf("%s is no a directory", name)
}
