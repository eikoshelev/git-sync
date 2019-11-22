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

// SSH or HTTP auth
func (g *git) gitAuth() {

	ep, err := transport.NewEndpoint(g.RemoteRepo)
	if err != nil {
		log.Fatalf("[Auth] Failed represents a Git URL in any supported protocol: %s\n", err.Error())
	}

	if strings.HasPrefix(ep.Protocol, "ssh") {
		keyFile, err := os.Open(g.SSHkey)
		if err != nil {
			log.Fatalf("[Auth] Failed open SSH key file: %s\n", err.Error())
		}
		
		defer keyFile.Close()

		key, err := ioutil.ReadAll(keyFile)
		if err != nil {
			log.Fatalf("[Auth] Failed read SSH key: %s\n", err.Error())
		}

		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatalf("[Auth] Failed parse SSH key: %s\n", err.Error())
		}

		g.Auth = &gitssh.PublicKeys{
			User:   "git",
			Signer: signer,
		}
	} else {
		g.Auth = &githttp.BasicAuth{
			Username: g.Login,
			Password: g.Password,
		}
	}
}
