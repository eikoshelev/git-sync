package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
)

// Data - flags structure
type Data struct {
	RemoteRepo string
	LocalPath  string
	CheckTime  int64
	SSHkey     string
	Login      string
	Password   string
	Force      bool
	Branch     string
}

func main() {

	var flags Data
	var err error
	var time string
	var force string

	flag.StringVar(&flags.RemoteRepo, "repo", os.Getenv("GIT_SYNC_REPO"), "URL to remote repository")
	flag.StringVar(&flags.LocalPath, "dir", os.Getenv("GIT_SYNC_ROOT"), "Path to local directory for repository")
	flag.StringVar(&time, "timer", os.Getenv("GIT_SYNC_WAIT"), "Timeout for check update (seconds)")
	flag.StringVar(&flags.SSHkey, "key", os.Getenv("GIT_SSH_KEY_PATH"), "Path to private ssh key for auth to the remote repository")
	flag.StringVar(&flags.Login, "login", os.Getenv("GIT_HTTP_LOGIN"), "Login for HTTP auth to the remote repository")
	flag.StringVar(&flags.Password, "pass", os.Getenv("GIT_HTTP_PASSWORD"), "Password for HTTP auth to the remote repository")
	flag.StringVar(&force, "force", os.Getenv("GIT_FORCE_PULL"), "Forced pull with changed local repository")
	flag.StringVar(&flags.Branch, "branch", os.Getenv("GIT_SYNC_BRANCH"), "Remote branch for pull")

	flag.Parse()

	if time != "" {
		flags.CheckTime, err = strconv.ParseInt(time, 10, 64)
		if err != nil {
			log.Printf("Failed convert 'GIT_SYNC_WAIT' (string) to int64: %v, set default timeout - 60s\n", err)
			flags.CheckTime = 60
		}
	} else {
		log.Printf("Flag '-timer' and/or 'GIT_SYNC_WAIT' env is empty - set default timeout - 60s\n")
	}

	if force != "" {
		flags.Force, err = strconv.ParseBool(force)
		if err != nil {
			log.Printf("Failed convert 'GIT_FORCE_PULL' (string) to bool: %v, set default rule - false\n", err)
		}
	} else {
		log.Printf("Flag '-force' and/or 'GIT_FORCE_PULL' env is empty - set default rule - false\n")
	}

	if flags.Branch == "" {
		flags.Branch = "master"
	}

	log.Printf("Started work!\nSet parameters:\nRemote repository: '%s'\nRemote branch: '%s'\nLocal directory: '%s'\nFetch timeout: %v seconds\nForce pull: %v\n\n", flags.RemoteRepo, flags.Branch, flags.LocalPath, flags.CheckTime, flags.Force)

	gitClone(flags)

	sgnl := make(chan os.Signal)
	signal.Notify(sgnl, os.Interrupt, os.Kill)
	<-sgnl
}
