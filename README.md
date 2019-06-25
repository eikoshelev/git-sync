[![Build Status](https://travis-ci.org/eikoshelev/git-sync.svg?branch=master)](https://travis-ci.org/eikoshelev/git-sync)

# git-sync

### Description
  
Real-time local repository update (every n seconds):
* When starting, it checks the presence of the repository at the specified path (if the repository already exists, it starts checking for updates, if not, it will first clone);
* Checks repository updates every n seconds;
* As soon as the remote repository is updated - makes "git pull" in the specified directory.

### Getting
```
git clone https://github.com/eikoshelev/git-sync.git
```
### Building
```
cd git-sync && go build
```
### Usage
```
./git-sync -repo <URL to remote repository> \
           -branch <branch of the monitored repository (default - master)>
           -dir <path to local directory for repository> \
           -key <path to private ssh key for auth to the remote repository (SSH protocol)> \
           -login <login for HTTP auth to the remote repository (HTTP protocol)> \
           -pass <password for HTTP auth to the remote repository (HTTP protocol)> \
           -timer <timeout for check update (seconds)> \
           -force <force pool if local repository is changed (default - false)>
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
|`GIT_FORCE_PULL`  | -force |
|`GIT_SYNC_BRANCH`  | -branch |
  
* Description of the flags used: ```./git-sync -h```
