package git

import (
	"github.com/eikoshelev/git-sync/internal/repository"

	"github.com/go-git/go-git/v5/plumbing/transport"
	"go.uber.org/zap"
)

type Git struct {
	SSHKeyPath, Login, Password string
	AuthMethod                  transport.AuthMethod
	Repository                  *repository.Repository
	Logger                      *zap.Logger
}

func Init(sshKeyPath, login, password string, repository *repository.Repository, logger *zap.Logger) *Git {
	return &Git{
		SSHKeyPath: sshKeyPath,
		Login:      login,
		Password:   password,
		Repository: repository,
		Logger:     logger,
	}
}
