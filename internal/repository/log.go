package repository

type commitInfo struct {
	Hash    string
	Author  string
	Email   string
	Time    string
	Message string
}

func (repo *Repository) GetLastCommit() (*commitInfo, error) {
	ref, err := repo.Local.Head()
	if err != nil {
		return nil, err
	}
	commit, err := repo.Local.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}
	return &commitInfo{
		Hash:    commit.Hash.String(),
		Author:  commit.Author.Name,
		Email:   commit.Author.Email,
		Time:    commit.Author.When.UTC().String(),
		Message: commit.Message,
	}, nil
}
