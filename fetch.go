package main

import (
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

// fetchCheck - chek update repository
func gitFetch(repos *git.Repository, url, key string) {

	skey, err := gitAuth(url, key)
	if err != nil {
		log.Fatalf("[Fetch Auth] Failed auth: %s", err)
	}
	err = repos.Fetch(&git.FetchOptions{
		Auth:     skey,
		Progress: os.Stdout,
		Force:    true,
	})
	if err != git.NoErrAlreadyUpToDate {
		log.Println("[Fetch] Update detected, pulling... ")
		gitPull(repos, url, key)
	}
}
