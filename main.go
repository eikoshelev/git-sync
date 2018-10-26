package main

import (
	"flag"
	"os"
	"os/signal"
)

func main() {

	remoteRepo := flag.String("repo", "", "URL to remote repository")
	localPath := flag.String("path", ".", "Path to local directory for repository")
	checkTime := flag.Int64("time", 60, "Timeout for check update (seconds)")
	sshKey := flag.String("key", "~/.ssh/id_rsa", "Path to private ssh key for auth to the remote repository")

	flag.Parse()

	gitClone(*sshKey, *remoteRepo, *localPath, *checkTime)

	sgnl := make(chan os.Signal)
	signal.Notify(sgnl, os.Interrupt, os.Kill)
	<-sgnl
}
