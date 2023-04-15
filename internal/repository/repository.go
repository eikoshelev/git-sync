package repository

import (
	"time"

	go_git "github.com/go-git/go-git/v5"
)

type Repository struct {
	URL, Tag, Branch, LocalPath string
	FetchTimeout                time.Duration
	ForcePull                   bool
	Local                       *go_git.Repository
}

func Init(url, tag, branch, localPath string, fetchTimeout time.Duration, forcePull bool) *Repository {
	return &Repository{
		URL:          url,
		Tag:          tag,
		LocalPath:    localPath,
		FetchTimeout: fetchTimeout,
		ForcePull:    forcePull,
	}
}
