package main

import (
	"log"

	go_git "gopkg.in/src-d/go-git.v4"
)

// chek update repository
func (g *git) gitFetch() {

	err := g.Repo.Fetch(&go_git.FetchOptions{
		Auth: g.Auth,
		//Force: true,
		//Progress: os.Stdout,
	})

	if err != go_git.NoErrAlreadyUpToDate {
		log.Printf("[Fetch] Update detected, pulling..\n")
		g.gitPull()
	} else if err != go_git.NoErrAlreadyUpToDate && err != nil {
		log.Printf("[Fetch] Error: %s\n", err.Error())
	}
}
