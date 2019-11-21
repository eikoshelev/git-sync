package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	go_git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

type git struct {
	RemoteRepo, Tag, Branch, LocalPath, SSHkey, Login, Password string
	CheckTime                                                   time.Duration
	Force                                                       bool

	// for methods:
	Auth transport.AuthMethod
	Repo *go_git.Repository
}

func main() {

	var g git

	if ok := g.checkParams(); ok {
		log.Printf("Set parameters:\nRemote repository: '%s'\nRemote tag: '%s'\nRemote branch: '%s'\nLocal directory: '%s'\nFetch timeout: %v\nForce pull: %v\n[HTTP login/password and SSH key path is hidden]\n\n", g.RemoteRepo, g.Tag, g.Branch, g.LocalPath, g.CheckTime, g.Force)
		go g.gitClone()
	} else {
		os.Exit(1)
	}

	sgnl := make(chan os.Signal)
	signal.Notify(sgnl, os.Interrupt, os.Kill)
	<-sgnl
}

func (g *git) checkParams() bool {

	var (
		err   error
		timer string
		force string
	)

	flag.StringVar(&g.RemoteRepo, "repo", os.Getenv("GIT_SYNC_REPO"), "URL to remote repository")
	flag.StringVar(&g.LocalPath, "dir", os.Getenv("GIT_SYNC_ROOT"), "Path to local directory for repository")
	flag.StringVar(&timer, "timer", os.Getenv("GIT_SYNC_WAIT"), "Timeout for check update, default — 1m")
	flag.StringVar(&g.SSHkey, "key", os.Getenv("GIT_SSH_KEY_PATH"), "Path to private SSH key for auth to the remote repository")
	flag.StringVar(&g.Login, "login", os.Getenv("GIT_HTTP_LOGIN"), "Login for HTTP auth to the remote repository")
	flag.StringVar(&g.Password, "pass", os.Getenv("GIT_HTTP_PASSWORD"), "Password for HTTP auth to the remote repository")
	flag.StringVar(&force, "force", os.Getenv("GIT_FORCE_PULL"), "Forced pull with changed local repository")
	flag.StringVar(&g.Tag, "tag", os.Getenv("GIT_SYNC_TAG"), "Remote tag for pull. NOTE (!): If the tag flag/env is specified, the specified branch flag/env will be ignored")
	flag.StringVar(&g.Branch, "branch", os.Getenv("GIT_SYNC_BRANCH"), "Remote branch for pull. NOTE (!): If the tag flag/env is specified, the specified branch flag/env will be ignored")

	flag.Parse()

	log.Printf("Checking parameters..")

	// for all checks
	checker := true

	// timeout for fetch
	if timer == "" {
		log.Printf("Flag '-timer' and/or 'GIT_SYNC_WAIT' env is empty: set default timeout - 60s\n")
		g.CheckTime = time.Duration(60 * time.Second)
	} else {
		g.CheckTime, err = time.ParseDuration(timer)
		if err != nil {
			log.Printf("Failed parse value of flag '-timer' and/or 'GIT_SYNC_WAIT' env: %s, set default timeout - 60s\n", err.Error())
			g.CheckTime = time.Duration(60 * time.Second)
		}
	}

	// force pull rule
	if force == "" {
		log.Printf("Flag '-force' and/or 'GIT_FORCE_PULL' env is empty: set default rule - false\n")
	} else {
		g.Force, err = strconv.ParseBool(force)
		if err != nil {
			log.Printf("Failed parse value of flag '-force' and/or 'GIT_FORCE_PULL' env: %s, set default rule - false\n", err.Error())
		}
	}

	// remote repository
	if g.RemoteRepo == "" {
		log.Printf("Flag '-repo' and/or 'GIT_SYNC_REPO' env is empty: specify which repository should be monitored\n")
		checker = false
	}

	// local directory
	if g.LocalPath == "" {
		log.Printf("Flag '-dir' and/or 'GIT_SYNC_ROOT' env is empty: specify in which directory you want to pull updates of the remote repository\n")
		checker = false
	}

	// remote branch or tag
	if g.Tag != "" {
		g.Branch = fmt.Sprintf("refs/tags/%s", g.Tag)
	} else {
		log.Printf("Flag '-tag' and/or 'GIT_SYNC_TAG' env is empty: working with branch\n")
		if g.Branch != "" {
			g.Branch = fmt.Sprintf("refs/heads/%s", g.Branch)
		} else {
			log.Printf("Flag '-branch' and/or 'GIT_SYNC_BRANCH' env is empty: set default branch - 'master'\n")
			g.Branch = "refs/heads/master"
		}
	}

	return checker
}
