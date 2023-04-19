package git

import (
	"os"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/transport"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

// SSH or HTTP auth
func (g *Git) Auth() error {
	ep, err := transport.NewEndpoint(g.Repository.URL)
	if err != nil {
		return err
	}
	if strings.HasPrefix(ep.Protocol, "ssh") {
		key, err := os.ReadFile(g.SSHKeyPath)
		if err != nil {
			return err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return err
		}
		g.AuthMethod = &gitssh.PublicKeys{
			User:   gitssh.DefaultUsername,
			Signer: signer,
		}
	} else {
		g.AuthMethod = &githttp.BasicAuth{
			Username: g.Login,
			Password: g.Password,
		}
	}
	return nil
}
