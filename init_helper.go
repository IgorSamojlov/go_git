package main

func repoInit() error {
	err := mkdir(REPO_DIR)
	if err != nil {
		return err
	}
	err = mkdir(REPO_DIR, OBJ_DIR)
	if err != nil {
		return err
	}
	return nil
}
