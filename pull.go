package main

import (
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

// pullRepo - pulling remote repository if there is an update
func gitPull(repo *git.Repository, data Data) {

	auth, err := gitAuth(data.RemoteRepo, data.SSHkey, data.Login, data.Password)
	if err != nil {
		log.Fatalf("[Pull Auth] Failed auth: %s", err)
	}

	wTree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("[Pull] Failed get work tree: %s", err)
	}

	err = wTree.Pull(&git.PullOptions{
		//SingleBranch:      true,
		Auth: auth,
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
