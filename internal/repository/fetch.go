package repository

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

// check update repository
func (repo *Repository) Fetch(auth transport.AuthMethod) error {
	return repo.Local.Fetch(&git.FetchOptions{
		Auth: auth,
	})
}
