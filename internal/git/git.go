package git

import (
	"os"
	"strings"

	"github.com/eikoshelev/git-sync/internal/repository"
	"go.uber.org/zap"

	"github.com/go-git/go-git/v5/plumbing/transport"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

type Git struct {
	SSHKeyPath, Login, Password string
	AuthMethod                  transport.AuthMethod
	Repository                  *repository.Repository
	Logger                      *zap.Logger
}

const SSHGitUser = "git"

func Init(sshKeyPath, login, password string, repository *repository.Repository, logger *zap.Logger) *Git {
	return &Git{
		SSHKeyPath: sshKeyPath,
		Login:      login,
		Password:   password,
		Repository: repository,
		Logger:     logger,
	}
}

// SSH or HTTP auth
func (git *Git) Auth() error {
	ep, err := transport.NewEndpoint(git.Repository.URL)
	if err != nil {
		return err
	}
	if strings.HasPrefix(ep.Protocol, "ssh") {
		key, err := os.ReadFile(git.SSHKeyPath)
		if err != nil {
			return err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return err
		}
		git.AuthMethod = &gitssh.PublicKeys{
			User:   SSHGitUser,
			Signer: signer,
		}
	} else {
		git.AuthMethod = &githttp.BasicAuth{
			Username: git.Login,
			Password: git.Password,
		}
	}
	return nil
}
