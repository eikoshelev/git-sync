package git

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (g *Git) Sync() {
	// authentication
	if err := g.Auth(); err != nil {
		g.Logger.Fatal("authentication failed", zap.String("reason", err.Error()))
	}
	g.Logger.Info("authentication success")
	// initial clone (or opening a local repository)
	err := g.Repository.Clone(g.AuthMethod)
	g.Logger.Info("debug", zap.Any("err clone", err))
	if err != nil {
		g.Logger.Fatal("clone repository failed", zap.String("reason", err.Error()))
	}
	g.PrintUpdate()
	// fetching update repo
	for {
		err := g.Repository.Fetch(g.AuthMethod)
		g.Logger.Info("debug", zap.Any("err fetch", err))
		if err == nil {
			g.Logger.Info("repository updates detected, try to pull")
			err := g.Repository.Pull(g.AuthMethod)
			g.Logger.Info("debug", zap.Any("err pull", err))
			if err != nil {
				g.Logger.Error("pull repository failed", zap.String("reason", err.Error()))
			}
			g.PrintUpdate()
		} else if !errors.Is(err, git.NoErrAlreadyUpToDate) {
			g.Logger.Error("fetch repository failed", zap.String("reason", err.Error()))
		}
		time.Sleep(g.Repository.FetchTimeout)
	}
}

func (g *Git) PrintUpdate() {
	// show last commit
	ref, err := g.Repository.Local.Head()
	if err != nil {
		g.Logger.Error("failed get the link pointed to by HEAD", zap.String("reason", err.Error()))
	}
	commit, err := g.Repository.Local.CommitObject(ref.Hash())
	if err != nil {
		g.Logger.Error("failed to get latest commit", zap.String("reason", err.Error()))
	}
	g.Logger.Info(
		"latest commit info",
		zap.String("hash", commit.Hash.String()),
		zap.String("author", commit.Author.Name),
		zap.String("email", commit.Author.Email),
		zap.String("time", commit.Author.When.String()),
		zap.String("message", commit.Message),
	)
}
