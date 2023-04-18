package repository

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/pkg/errors"
)

func (repo *Repository) Clone(auth transport.AuthMethod) error {
	var err error
	// checking for the existence of a local repository
	repo.Local, err = git.PlainOpen(repo.LocalPath)
	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			repo.Local, err = git.PlainClone(repo.LocalPath, false, &git.CloneOptions{
				URL:           repo.URL,
				Auth:          auth,
				ReferenceName: plumbing.ReferenceName(repo.Branch),
				SingleBranch:  true,
			})
			if err != nil {
				return errors.Wrap(err, "clone repository failed")
			}
			return nil
		}
		return errors.Wrap(err, "open local repository failed")
	}
	return nil
}
