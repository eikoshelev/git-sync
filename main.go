package main

import (
	"flag"
	"os"
	"os/signal"
)

func main() {

	remoteRepo := flag.String("repo", "", "URL to remote repository")
	localPath := flag.String("path", ".", "Local directory for repository")
	//	sshKey := flag.String()

	flag.Parse()

	gitClone(*remoteRepo, *localPath)

	sgnl := make(chan os.Signal)
	signal.Notify(sgnl, os.Interrupt, os.Kill)
	<-sgnl
}
