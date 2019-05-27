package main

import (
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// pullRepo - pulling remote repository if there is an update
func gitPull(repo *git.Repository, data Data) {

	auth, err := gitAuth(data.RemoteRepo, data.SSHkey, data.Login, data.Password)
	if err != nil {
		log.Printf("[Pull Auth] Failed auth: %v\n", err)
	}

	wTree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("[Pull] Failed get work tree: %v\n", err)
	}

	err = wTree.Pull(&git.PullOptions{
		ReferenceName: plumbing.ReferenceName("refs/heads/" + data.Branch),
		Auth:          auth,
		SingleBranch:  true,
		Progress:      os.Stdout,
		Force:         true,
	})

	if err == nil {
		log.Printf("[Pull] Success!\n")
	} else {
		switch err {
		case git.ErrUnstagedChanges:
			if data.Force != true {
				log.Printf("[Pull] Error: %v (local repository changed). Rule for forced pull - %v. Can`t force pull.\n", err, data.Force)
			} else {
				log.Printf("[Pull] Info: %v (local repository changed). Rule for forced pull - %v\n. Force pulling...", err, data.Force)
				err := wTree.Reset(&git.ResetOptions{
					Mode: git.ResetMode(1),
				})
				if err == nil {
					log.Printf("[Pull] Success!\n")
				} else {
					log.Printf("[Pull] Error: %v\n", err)
				}
			}
		case git.NoErrAlreadyUpToDate:
			log.Printf("[Pull] Nothing to pull: %v\n", err)
		default:
			log.Printf("[Pull] Error: %v\n", err)
		}
	}

	// show last commit
	ref, err := repo.Head()
	if err != nil {
		log.Printf("Failed get reference where HEAD is pointing to: %v\n", err)
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		log.Printf("[Pull] Can`t show last commit: %v\n", err)
	} else {
		log.Printf("[Pull] Last commit: %v\n", commit)
	}
}
