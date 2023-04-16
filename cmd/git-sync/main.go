package main

import (
	"github.com/eikoshelev/git-sync/internal/config"
	"github.com/eikoshelev/git-sync/internal/git"
	"github.com/eikoshelev/git-sync/internal/logger"
	"github.com/eikoshelev/git-sync/internal/repository"
	"go.uber.org/zap"
)

func main() {
	logger := logger.GetLogger()
	config, err := config.GetConfiguration()
	if err != nil {
		logger.Fatal("configuration failed", zap.String("reason", err.Error()))
	}
	logger.Info("configuration success", zap.Any("config", config))
	repository := repository.Init(
		config.RemoteRepoURL,
		config.RemoteRepoTag,
		config.RemoteRepoBranch,
		config.LocalRepoPath,
		config.FetchTimeout,
		config.ForcePull,
	)
	git := git.Init(
		config.SSHKeyPath,
		config.HTTPLogin,
		config.HTTPPassword,
		repository,
		&logger,
	)
	git.Sync()
}
