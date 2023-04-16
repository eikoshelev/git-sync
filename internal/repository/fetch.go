package repository

import (
	go_git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

// check update repository
func (repo *Repository) Fetch(auth transport.AuthMethod) error {
	return repo.Local.Fetch(&go_git.FetchOptions{
		Auth: auth,
		//Force: true,
		//Progress: os.Stdout,
	})
}
