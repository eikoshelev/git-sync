package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

// gitAuth - ssh/http auth
func gitAuth(uri string, key string) (transport.AuthMethod, error) {

	var auth transport.AuthMethod

	ep, err := transport.NewEndpoint(uri)
	if err != nil {
		return nil, err
	}

	// ssh auth
	if strings.HasPrefix(ep.Protocol, "ssh") && key != "" {
		var signer ssh.Signer

		sshFile, err := os.Open(key)
		if err != nil {
			log.Fatalf("[Auth] Failed open SSH key file: %s", err)
		}
		sshB, err := ioutil.ReadAll(sshFile)
		if err != nil {
			log.Fatalf("[Auth] Failed read SSH key: %s", err)
		}

		signer, err = ssh.ParsePrivateKey(sshB)
		if err != nil {
			log.Fatalf("[Auth] Failed parse SSH key: %s", err)
		}

		sshAuth := &gitssh.PublicKeys{User: "git", Signer: signer}
		return sshAuth, nil
	}

	// http auth
	if strings.HasPrefix(ep.Protocol, "http") && ep.User != "" && ep.Password != "" {
		auth = &githttp.BasicAuth{
			Username: ep.User,
			Password: ep.Password,
		}
	}

	return auth, nil
}
