package main

import (
	"log"
	"os"
	"time"

	git "gopkg.in/src-d/go-git.v4"
)

// cloneRepo - cloning remote repository
func gitClone(key, url, dir string, timer int64) {

	skey, err := gitAuth(url, key)
	if err != nil {
		log.Fatalf("[Clone Auth] Failed auth: %s", err)
	}

	// check repository availability in local directory
	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Printf("[Clone] Repository not found in '%s' cloning...", dir)
		repo, err = git.PlainClone(dir, false, &git.CloneOptions{
			URL:  url,
			Auth: skey,
			//SingleBranch:      true,
			//RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatalf("[Clone] Failed clone remote repository: %s", err)
		}
	} else {
		log.Printf("Repository found in '%s' opening...", dir)
	}

	go func() {
		for {
			gitFetch(repo, url, key)
			time.Sleep(time.Duration(timer) * time.Second)
		}
	}()
}
