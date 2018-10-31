package main

import (
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

// fetchCheck - chek update repository
func gitFetch(repo *git.Repository, data Data) {

	auth, err := gitAuth(data.RemoteRepo, data.SSHkey, data.Login, data.Password)
	if err != nil {
		log.Fatalf("[Fetch Auth] Failed auth: %s", err)
	}
	err = repo.Fetch(&git.FetchOptions{
		Auth:     auth,
		Progress: os.Stdout,
		Force:    true,
	})
	if err != git.NoErrAlreadyUpToDate {
		log.Println("[Fetch] Update detected, pulling... ")
		gitPull(repo, data)
	}
}
