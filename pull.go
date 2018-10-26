package main

import (
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

// pullRepo - pulling remote repository if there is an update
func gitPull(repository *git.Repository, url, key string) {

	skey, err := gitAuth(url, key)
	if err != nil {
		log.Fatalf("[Pull Auth] Failed auth: %s", err)
	}
	wTree, err := repository.Worktree()
	if err != nil {
		log.Fatalf("[Pull] Failed get work tree: %s", err)
	}

	err = wTree.Pull(&git.PullOptions{
		//SingleBranch:      true,
		Auth: skey,
		//RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress: os.Stdout,
		Force:    true,
	})
	if err != nil {
		log.Fatalf("[Pull] Failed pull remote repository: %s", err)
	} else {
		log.Println("[Pull] Success!")
	}
}
