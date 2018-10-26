# git-sync

### Description

Update local repository in real time.
1. At startup it makes cloning of the remote repository;
2. After a certain period of time checks for updates in the remote repository;
3. In the presence of updates makes "git pull".
  
### Getting

```
git clone https://github.com/eikoshelev/git-sync
``` 
```
cd git-sync
```  
```
go build
```
  
### Usage

```
./git-sync -repo="<url to remote repository>" -dir="<path to directory for clone repository>" -key="<path to private ssh key for auth to the remote repository>" -time=<timeout for check update (seconds, default - 60 sec)>
```  
More information: "git-sync -h"
```
$ ./git-sync -h
Usage of ./git-sync:
  -dir string
    	Path to local directory for repository (default ".")
  -key string
    	Path to private ssh key for auth to the remote repository (default "~/.ssh/id_rsa")
  -repo string
    	URL to remote repository
  -time int
    	Timeout for check update (seconds) (default 60)
```
