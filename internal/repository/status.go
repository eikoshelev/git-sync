package repository

import "github.com/pkg/errors"

func (repo *Repository) Status() (bool, error) {
	wt, err := repo.Local.Worktree()
	if err != nil {
		return false, errors.Wrap(err, "failed to checkout working branch")
	}
	st, err := wt.Status()
	if err != nil {
		return false, errors.Wrap(err, "failed to get work branch status")
	}
	return st.IsClean(), nil
}
