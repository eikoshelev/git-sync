package repository

import (
	"log"
	"os"

	go_git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

// pulling remote repository if there is an update
func (repo *Repository) Pull(auth transport.AuthMethod) error {
	wTree, err := repo.Local.Worktree()
	if err != nil {
		return err
	}

	err = wTree.Pull(&go_git.PullOptions{
		ReferenceName: plumbing.ReferenceName(repo.Branch),
		Auth:          auth,
		SingleBranch:  true,
		Progress:      os.Stdout,
		//Force:         true,
	})

	if err != nil {
		switch err {
		case go_git.ErrUnstagedChanges:
			if repo.ForcePull {
				err := wTree.Reset(&go_git.ResetOptions{
					Mode: go_git.ResetMode(1),
				})
				if err != nil {
					return err
				}
			} else {
				log.Printf("[Pull] Error: %s (local repository changed).\nRule for forced pull - %v. Can`t force pull.\n", err.Error(), repo.ForcePull)
			}
		case go_git.NoErrAlreadyUpToDate:
			log.Printf("[Pull] Nothing to pull: %s\n", err.Error())
		default:
			return err
		}
	}
	return nil
}
