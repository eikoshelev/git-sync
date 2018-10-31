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
func gitAuth(url, key, login, pass string) (transport.AuthMethod, error) {

	var auth transport.AuthMethod

	ep, err := transport.NewEndpoint(url)
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
	if strings.HasPrefix(ep.Protocol, "http") && login != "" && pass != "" {
		auth = &githttp.BasicAuth{
			Username: login,
			Password: pass,
		}
	} else {
		log.Fatalf("Missing login and/or password for HTTP auth: %s", err)
	}

	return auth, nil
}
