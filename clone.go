package main

import (
	"log"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// cloneRepo - cloning remote repository
func gitClone(data Data) {

	auth, err := gitAuth(data.RemoteRepo, data.SSHkey, data.Login, data.Password)
	if err != nil {
		log.Fatalf("Failed auth: %s", err)
	}

	//check repository availability in local directory
	repo, err := git.PlainOpen(data.LocalPath)
	if err != nil {
		log.Printf("[Clone] Repository not found in '%s', cloning...\n", data.LocalPath)

		repo, err = git.PlainClone(data.LocalPath, false, &git.CloneOptions{
			URL:           data.RemoteRepo,
			ReferenceName: plumbing.ReferenceName("refs/heads/" + data.Branch),
			Auth:          auth,
			SingleBranch:  true,
			//Progress: os.Stdout,
		})
		if err != nil {
			log.Fatalf("[Clone] Failed clone remote repository: %v\n", err)
		} else {
			log.Printf("[Clone] Success!\n")
		}

		// show last commit
		ref, err := repo.Head()
		if err != nil {
			log.Printf("Failed get reference where HEAD is pointing to: %v\n", err)
		}

		commit, err := repo.CommitObject(ref.Hash())
		if err != nil {
			log.Printf("[Clone] Can`t show last commit: %v\n", err)
		} else {
			log.Printf("[Clone] Last commit: %v\n", commit)
		}
	} else {
		log.Printf("Repository found in '%s' opening...\n", data.LocalPath)
	}

	for {
		gitFetch(repo, data)
		time.Sleep(time.Duration(data.CheckTime) * time.Second)
	}
}
