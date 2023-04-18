package repository

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/pkg/errors"
)

// pulling remote repository if there is an update
func (repo *Repository) Pull(auth transport.AuthMethod) error {
	wTree, err := repo.Local.Worktree()
	if err != nil {
		return err
	}
	err = wTree.Pull(&git.PullOptions{
		ReferenceName: plumbing.ReferenceName(repo.Branch),
		Auth:          auth,
		SingleBranch:  true,
	})
	if err != nil {
		if errors.Is(err, git.ErrUnstagedChanges) {
			if repo.ForcePull {
				if err := wTree.Reset(&git.ResetOptions{Mode: git.ResetMode(1)}); err != nil {
					return err
				}
				return nil
			}
			return errors.Wrap(err, "forced pull required for update, but GSYNC_FORCE_PULL option is 'false'")
		}
		return err
	}
	return nil
}
