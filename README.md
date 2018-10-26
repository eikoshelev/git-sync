# git-sync

### Description

Update local repository in real time.
* Check repository availability in local directory (if the repository is found - it checks for updates, if not - then it clones);
* After a certain period of time checks for updates in the remote repository;
* In the presence of updates makes "git pull".
  
### Getting

```
git clone https://github.com/eikoshelev/git-sync
``` 
```
cd git-sync
```  
```
vgo build
```
  
### Usage

```
./git-sync -repo="git@github.com:eikoshelev/git-sync.git" -dir="/tmp/git-sync" -key="~/.ssh/id_rsa" -timer=60
```  
* More information: ```./git-sync -h```
```
$ ./git-sync -h
Usage of ./git-sync:
  -dir string
    	Path to local directory for repository (default - ".")
  -key string
    	Path to private ssh key for auth to the remote repository (default - "~/.ssh/id_rsa")
  -repo string
    	URL to remote repository
  -timer int
    	Timeout for check update (seconds) (default - 60)
```
