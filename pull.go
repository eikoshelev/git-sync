package main

import (
	"log"
	"os"

	go_git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// pulling remote repository if there is an update
func (g *git) gitPull() {

	wTree, err := g.Repo.Worktree()
	if err != nil {
		log.Fatalf("[Pull] Failed get work tree: %s\n", err.Error())
	}

	err = wTree.Pull(&go_git.PullOptions{
		ReferenceName: plumbing.ReferenceName(g.Branch),
		Auth:          g.Auth,
		SingleBranch:  true,
		Progress:      os.Stdout,
		//Force:         true,
	})

	if err == nil {
		log.Printf("[Pull] Success!\n")
	} else {
		switch err {
		case go_git.ErrUnstagedChanges:
			if g.Force {
				log.Printf("[Pull] Info: %s (local repository changed).\nRule for forced pull - %v. Force pulling...", err.Error(), g.Force)
				err := wTree.Reset(&go_git.ResetOptions{
					Mode: go_git.ResetMode(1),
				})
				if err == nil {
					log.Printf("[Pull] Success!\n")
				} else {
					log.Printf("[Pull] Error: %s\n", err.Error())
				}
			} else {
				log.Printf("[Pull] Error: %s (local repository changed).\nRule for forced pull - %v. Can`t force pull.\n", err.Error(), g.Force)
			}
		case go_git.NoErrAlreadyUpToDate:
			log.Printf("[Pull] Nothing to pull: %s\n", err.Error())
		default:
			log.Printf("[Pull] Error: %s\n", err.Error())
		}
	}

	// show last commit
	ref, err := g.Repo.Head()
	if err != nil {
		log.Printf("Failed get reference where HEAD is pointing to: %s\n", err.Error())
	}
	commit, err := g.Repo.CommitObject(ref.Hash())
	if err != nil {
		log.Printf("[Pull] Can`t show last commit: %s\n", err.Error())
	} else {
		log.Printf("[Pull] Last commit: %s\n", commit)
	}
}
