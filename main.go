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
}

func main() {

	var flags Data
	var err error
	var time string

	flag.StringVar(&flags.RemoteRepo, "repo", os.Getenv("GIT_SYNC_REPO"), "URL to remote repository")
	flag.StringVar(&flags.LocalPath, "dir", os.Getenv("GIT_SYNC_ROOT"), "Path to local directory for repository")
	flag.StringVar(&time, "timer", os.Getenv("GIT_SYNC_WAIT"), "Timeout for check update (seconds)")
	flag.StringVar(&flags.SSHkey, "key", os.Getenv("GIT_SSH_KEY_PATH"), "Path to private ssh key for auth to the remote repository")
	flag.StringVar(&flags.Login, "login", os.Getenv("GIT_HTTP_LOGIN"), "Login for HTTP auth to the remote repository")
	flag.StringVar(&flags.Password, "pass", os.Getenv("GIT_HTTP_PASSWORD"), "Password for HTTP auth to the remote repository")

	flag.Parse()

	flags.CheckTime, err = strconv.ParseInt(time, 10, 64)
	if err != nil {
		log.Printf("Failed convert 'GIT_SYNC_WAIT' (string) to int64: %s, set default timeout - 60s", err)
		flags.CheckTime = 60
	}

	gitClone(flags)

	sgnl := make(chan os.Signal)
	signal.Notify(sgnl, os.Interrupt, os.Kill)
	<-sgnl
}
