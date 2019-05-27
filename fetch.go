package main

import (
	"log"

	git "gopkg.in/src-d/go-git.v4"
)

// fetchCheck - chek update repository
func gitFetch(repo *git.Repository, data Data) {

	auth, err := gitAuth(data.RemoteRepo, data.SSHkey, data.Login, data.Password)
	if err != nil {
		log.Fatalf("[Fetch] Failed auth: %v\n", err)
	}
	err = repo.Fetch(&git.FetchOptions{
		Auth:  auth,
		Force: true,
		//Progress: os.Stdout,
	})
	if err != git.NoErrAlreadyUpToDate {
		log.Printf("[Fetch] Update detected, pulling...\n")
		gitPull(repo, data)
	} else if err != git.NoErrAlreadyUpToDate && err != nil {
		log.Printf("[Fetch] Error: %v\n", err)
	}
}
