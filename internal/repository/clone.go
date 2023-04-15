package repository

import (
	"github.com/pkg/errors"

	go_git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

func (repo *Repository) Clone(auth transport.AuthMethod) error {
	var err error
	// checking for the existence of a local repository
	repo.Local, err = go_git.PlainOpen(repo.LocalPath)
	if err != nil {
		if errors.Is(err, go_git.ErrRepositoryNotExists) {
			repo.Local, err = go_git.PlainClone(repo.LocalPath, false, &go_git.CloneOptions{
				URL:           repo.URL,
				ReferenceName: plumbing.ReferenceName(repo.Branch),
				Auth:          auth,
				SingleBranch:  true,
				//Progress: os.Stdout,
			})
			if err != nil {
				return err
			} else {
				return nil
			}
		}
	}
	return nil
}
