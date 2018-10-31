# git-sync

### Description
  
Real-time local repository update (every n seconds):
* When starting, it checks the presence of the repository at the specified path (if the repository already exists, it starts checking for updates, if not, it will first clone);
* Checks repository updates every n seconds;
* As soon as the remote repository is updated - makes "git pull" in the specified directory.

### Usage

```
./git-sync -repo <url to remote repo> -dir <destination directory> -key <the path to the private ssh key for authentication (for SSH)> -timer <update check interval (60 seconds by default)> -login <login for HTTP> -pass <password for HTTP>
```  
* git-sync defaults to using environment variables (for docker/kubernetes) if the flags are not explicitly set at startup:

| env variable   | flag |
|:---------------|:------
|`GIT_SYNC_REPO` | -repo |
|`GIT_SYNC_ROOT` | -dir |
|`GIT_SYNC_WAIT` | -timer |
|`GIT_SSH_KEY_PATH` | -key |
|`SSH_KNOWN_HOSTS` | - |
|`GIT_HTTP_LOGIN` | -login |
|`GIT_HTTP_PASSWORD` | -pass |
---
* Description of the flags used: ```./git-sync -h```
```
$ ./git-sync -h
Usage of ./git-sync:
  -dir string
    	Path to local directory for repository
  -key string
    	Path to private ssh key for auth to the remote repository
  -login string
    	Login for HTTP auth to the remote repository
  -pass string
    	Password for HTTP auth to the remote repository
  -repo string
    	URL to remote repository
  -timer string
    	Timeout for check update (seconds)
```
