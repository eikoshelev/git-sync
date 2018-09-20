package main

import (
	"log"
	"os"
	"time"

	git "gopkg.in/src-d/go-git.v4"
)

// cloneRepo - cloning remote repository
func cloneRepo(url, dir string) {
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Printf("[cloneRepo] Error: %v", err)
	}

	go func() {
		for {
			fetchCheck(repo)
			time.Sleep(5 * time.Second)
		}
	}()
}

// fetchCheck - chek update repository
func fetchCheck(repos *git.Repository) {
	err := repos.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})
	if err == git.NoErrAlreadyUpToDate {
		log.Printf("[fetchCheck] Nothing update: %v", err)
	} else {
		log.Printf("[fetchCheck] There are updates, Downloading... : %v", err)
		pullRepo(repos)
	}
}

// pullRepo - pulling remote repository if there is an update
func pullRepo(repository *git.Repository) {
	wTree, err := repository.Worktree()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	err = wTree.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})
	if err != nil {
		log.Printf("[pullRepo] Error: %v", err)
	}
}
