package git

import (
	"errors"
	"log"
	"time"

	go_git "github.com/go-git/go-git/v5"
	"go.uber.org/zap"
)

func (git *Git) Sync() {
	// authentication
	if err := git.Auth(); err != nil {
		git.Logger.Fatal("authentication failed", zap.String("reason", err.Error()))
	}
	git.Logger.Info("authentication success")
	// initial clone (or opening a local repository)
	if err := git.Repository.Clone(git.AuthMethod); err != nil {
		git.Logger.Fatal("clone repository failed", zap.String("reason", err.Error()))
	}
	git.PrintUpdate()
	// fetching update repo
	for {
		if err := git.Repository.Fetch(git.AuthMethod); err != nil {
			if errors.Is(err, go_git.NoErrAlreadyUpToDate) {
				continue
			}
			log.Printf("[Fetch] Update detected, pulling..\n")
			if err := git.Repository.Pull(git.AuthMethod); err != nil {
				git.Logger.Error("pull failed", zap.String("reason", err.Error()))
			}
			git.PrintUpdate()
		} else if err != go_git.NoErrAlreadyUpToDate && err != nil {
			log.Printf("[Fetch] Error: %s\n", err.Error())
		}
		time.Sleep(git.Repository.FetchTimeout)
	}
}

func (git *Git) PrintUpdate() {
	// show last commit
	ref, err := git.Repository.Local.Head()
	if err != nil {
		log.Printf("Failed get reference where HEAD is pointing to: %s\n", err.Error())
	}
	commit, err := git.Repository.Local.CommitObject(ref.Hash())
	if err != nil {
		log.Printf("[Clone] Can`t show last commit: %s\n", err.Error())
	} else {
		log.Printf("[Clone] Last commit: %v\n", commit)
	}
}
