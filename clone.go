package main

import (
	"log"
	"os"
	"time"

	git "gopkg.in/src-d/go-git.v4"
)

// cloneRepo - cloning remote repository
func gitClone(data Data) {

	auth, err := gitAuth(data.RemoteRepo, data.SSHkey, data.Login, data.Password)
	if err != nil {
		log.Fatalf("[Clone Auth] Failed auth: %s", err)
	}

	// check repository availability in local directory
	repo, err := git.PlainOpen(data.LocalPath)
	if err != nil {
		log.Printf("Repository not found in '%s' cloning...", data.LocalPath)
		repo, err = git.PlainClone(data.LocalPath, false, &git.CloneOptions{
			URL:  data.RemoteRepo,
			Auth: auth,
			//SingleBranch:      true,
			//RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatalf("[Clone] Failed clone remote repository: %s", err)
		}
	} else {
		log.Printf("Repository found in '%s' opening...", data.LocalPath)
	}

	for {
		gitFetch(repo, data)
		time.Sleep(time.Duration(data.CheckTime) * time.Second)
	}
}
