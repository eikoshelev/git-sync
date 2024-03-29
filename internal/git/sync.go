package git

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (g *Git) Sync() {
	// authentication
	if err := g.Auth(); err != nil {
		g.Logger.Fatal("git authentication failed", zap.String("reason", err.Error()))
	}
	g.Logger.Info("git authentication successful")
	// initial clone (or opening a local repository)
	err := g.Repository.Clone(g.AuthMethod)
	if err != nil {
		g.Logger.Fatal("clone repository failed", zap.String("reason", err.Error()))
	}
	commit, err := g.Repository.GetLastCommit()
	if err != nil {
		g.Logger.Error("failed to get latest commit", zap.String("reason", err.Error()))
	}
	g.Logger.Info("latest commit info", zap.Any("commit", commit))
	ok, err := g.Repository.Status()
	if err != nil {
		g.Logger.Error("work branch status unknown", zap.String("reason", err.Error()))
	}
	if !ok && !g.Repository.ForcePull {
		msg := fmt.Sprintf("one or more files have modified status in %s, GSYNC_FORCE_PULL option is 'false', remote repository updates will not be pulled", g.Repository.LocalPath)
		g.Logger.Warn(msg)
	}
	// fetching update repo
	g.Logger.Info("start fetch updates")
	for {
		err := g.Repository.Fetch(g.AuthMethod)
		if err == nil {
			g.Logger.Info("repository updates detected")
			err := g.Repository.Pull(g.AuthMethod)
			if err != nil {
				g.Logger.Error("pull repository failed", zap.String("reason", err.Error()))
			} else {
				commit, err := g.Repository.GetLastCommit()
				if err != nil {
					g.Logger.Error("pull repository success, but failed to get latest commit", zap.String("reason", err.Error()))
				} else {
					g.Logger.Info("pull repository success", zap.Any("commit", commit))
				}
			}
		} else if errors.Is(err, git.ErrRemoteNotFound) {
			g.Logger.Fatal(err.Error())
		} else if !errors.Is(err, git.NoErrAlreadyUpToDate) {
			g.Logger.Error("fetch repository failed", zap.String("reason", err.Error()))
		}
		time.Sleep(g.Repository.FetchTimeout)
	}
}
