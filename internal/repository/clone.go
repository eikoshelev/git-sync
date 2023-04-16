package repository

import (
	"github.com/pkg/errors"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

func (repo *Repository) Clone(auth transport.AuthMethod) error {
	var err error
	// checking for the existence of a local repository
	repo.Local, err = git.PlainOpen(repo.LocalPath)
	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			repo.Local, err = git.PlainClone(repo.LocalPath, false, &git.CloneOptions{
				URL:           repo.URL,
				ReferenceName: plumbing.ReferenceName(repo.Branch),
				Auth:          auth,
				SingleBranch:  true,
			})
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}
