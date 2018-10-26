package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

// ssh/http auth
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
			log.Printf("[Auth] Failed open SSH key file: %s", err)
		}
		sshB, err := ioutil.ReadAll(sshFile)
		if err != nil {
			log.Printf("[Auth] Failed read SSH key: %s", err)
		}

		signer, err = ssh.ParsePrivateKey(sshB)
		if err != nil {
			log.Printf("[Auth] Failed parse SSH key: %s", err)
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

// cloneRepo - cloning remote repository
func gitClone(key, url, dir string, timer int64) {
	skey, err := gitAuth(url, key)
	if err != nil {
		log.Printf("[Clone Auth] Failed auth: %s", err)
	}
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               url,
		Auth:              skey,
		SingleBranch:      true,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})
	if err != nil {
		log.Printf("[Clone] Failed clone remote repository: %s", err)
	}

	go func() {
		for {
			gitFetch(repo, url, key)
			time.Sleep(time.Duration(timer) * time.Second)
		}
	}()
}

// fetchCheck - chek update repository
func gitFetch(repos *git.Repository, url, key string) {
	skey, err := gitAuth(url, key)
	if err != nil {
		log.Printf("[Fetch Auth] Failed auth: %s", err)
	}
	err = repos.Fetch(&git.FetchOptions{
		Auth:     skey,
		Progress: os.Stdout,
		Force:    true,
	})
	if err != git.NoErrAlreadyUpToDate {
		log.Println("[Fetch] Update detected, pulling.. ")
		gitPull(repos, url, key)
	}
}

// pullRepo - pulling remote repository if there is an update
func gitPull(repository *git.Repository, url, key string) {
	skey, err := gitAuth(url, key)
	if err != nil {
		log.Printf("[Pull Auth] Failed auth: %s", err)
	}
	wTree, err := repository.Worktree()
	if err != nil {
		log.Printf("[Pull] Failed get work tree: %s", err)
	}

	err = wTree.Pull(&git.PullOptions{
		SingleBranch:      true,
		Auth:              skey,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
		Force:             true,
	})
	if err != nil {
		log.Printf("[Pull] Failed pull remote repository: %s", err)
	} else {
		log.Println("[Pull] Success!")
	}
}
