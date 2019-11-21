package main

import (
	"log"
	"time"

	go_git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// cloning remote repository
func (g *git) gitClone() {

	g.gitAuth()

	log.Printf("Authentication successful! Starting work..\n")

	var err error

	// check repository availability in local directory
	g.Repo, err = go_git.PlainOpen(g.LocalPath)
	if err != nil {
		log.Printf("[Clone] Repository not found in '%s', cloning..\n", g.LocalPath)

		g.Repo, err = go_git.PlainClone(g.LocalPath, false, &go_git.CloneOptions{
			URL:           g.RemoteRepo,
			ReferenceName: plumbing.ReferenceName(g.Branch),
			Auth:          g.Auth,
			SingleBranch:  true,
			//Progress: os.Stdout,
		})
		if err != nil {
			log.Fatalf("[Clone] Failed clone remote repository: %s\n", err.Error())
		} else {
			log.Printf("[Clone] Success!\n")
		}

		// show last commit
		ref, err := g.Repo.Head()
		if err != nil {
			log.Printf("Failed get reference where HEAD is pointing to: %s\n", err.Error())
		}

		commit, err := g.Repo.CommitObject(ref.Hash())
		if err != nil {
			log.Printf("[Clone] Can`t show last commit: %s\n", err.Error())
		} else {
			log.Printf("[Clone] Last commit: %v\n", commit)
		}
	} else {
		log.Printf("Repository found in '%s' opening..\n", g.LocalPath)
	}

	// fetching update repo
	for {
		g.gitFetch()
		time.Sleep(g.CheckTime)
	}
}
