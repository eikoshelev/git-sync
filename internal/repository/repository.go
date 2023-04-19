package repository

import (
	"time"

	"github.com/go-git/go-git/v5"
)

type Repository struct {
	URL, Tag, Branch, LocalPath string
	FetchTimeout                time.Duration
	ForcePull                   bool
	Local                       *git.Repository
}

func Init(url, tag, branch, localPath string, fetchTimeout time.Duration, forcePull bool) *Repository {
	return &Repository{
		URL:          url,
		Tag:          tag,
		Branch:       branch,
		LocalPath:    localPath,
		FetchTimeout: fetchTimeout,
		ForcePull:    forcePull,
	}
}
