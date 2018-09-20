package main

import (
	"log"
	"os"
	"time"

	git "gopkg.in/src-d/go-git.v4"
)

// cloneRepo - cloning remote repository
func gitClone(url, dir string) {
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Printf("[Clone] Failed clone remote repository: %v", err)
	}

	go func() {
		for {
			gitFetch(repo)
			time.Sleep(5 * time.Second)
		}
	}()
}

// fetchCheck - chek update repository
func gitFetch(repos *git.Repository) {
	err := repos.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})
	if err == git.NoErrAlreadyUpToDate {
		log.Printf("[Fetch] Nothing update: %v", err)
	} else {
		log.Printf("[Fetch] Update detected, Downloading.. : %v", err)
		gitPull(repos)
	}
}

// pullRepo - pulling remote repository if there is an update
func gitPull(repository *git.Repository) {
	wTree, err := repository.Worktree()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	err = wTree.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})
	if err != nil {
		log.Printf("[Pull] Failed pull remote repository: %v", err)
	}
}
