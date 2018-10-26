package main

import (
	"errors"
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
			return nil, errors.New("[Auth] Failed open SSH key file: " + err.Error())
		}
		sshB, err := ioutil.ReadAll(sshFile)
		if err != nil {
			return nil, errors.New("[Auth] Failed read SSH key: " + err.Error())
		}

		signer, err = ssh.ParsePrivateKey(sshB)
		if err != nil {
			return nil, errors.New("[Auth] Failed parse SSH key: " + err.Error())
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
		log.Printf("[Auth] Failed auth: " + err.Error())
	}
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,
		Auth:     skey,
		Depth:    1,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Printf("[Clone] Failed clone remote repository: " + err.Error())
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
		log.Printf("[Auth] Failed auth" + err.Error())
	}
	err = repos.Fetch(&git.FetchOptions{
		Depth:    1,
		Auth:     skey,
		Progress: os.Stdout,
	})
	if err == git.NoErrAlreadyUpToDate {
		log.Printf("[Fetch] Nothing update: " + err.Error())
	} else {
		log.Printf("[Fetch] Update detected, Downloading.. : " + err.Error())
		gitPull(repos, url, key)
	}
}

// pullRepo - pulling remote repository if there is an update
func gitPull(repository *git.Repository, url, key string) {
	skey, err := gitAuth(url, key)
	if err != nil {
		log.Printf("[Auth] Failed auth" + err.Error())
	}
	wTree, err := repository.Worktree()
	if err != nil {
		log.Printf("[Pull] Failed get work tree: " + err.Error())
	}

	err = wTree.Pull(&git.PullOptions{
		Auth:     skey,
		Progress: os.Stdout,
		Force:    true,
	})
	if err != nil {
		log.Printf("[Pull] Failed pull remote repository: ", err.Error())
	}
}
