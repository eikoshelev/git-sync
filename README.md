[![Build Status](https://travis-ci.org/eikoshelev/git-sync.svg?branch=master)](https://travis-ci.org/eikoshelev/git-sync)

# git-sync

:recycle: Easy to use and stable update of your repository.

### Description
  
Real-time local repository update (every N time):
* When starting, it checks the presence of the repository at the specified path (if the repository already exists, it starts checking for updates, if not, it will first clone);
* Checks repository updates every N time;
* As soon as the remote repository is updated - makes `git pull` in the specified directory.

### Getting
```
git clone https://github.com/eikoshelev/git-sync.git
```
### Building
```
cd git-sync && go build
```
### Usage

* Description of the flags used: ```./git-sync -h```

```
./git-sync -h
Usage of ./git-sync:
  -branch string
    	Remote branch for pull. NOTE (!): If the 'tag' flag/env is specified, the 'branch' flag/env will be ignored
  -dir string
    	Path to local directory for repository
  -force string
    	Forced pull with changed local repository
  -key string
    	Path to private SSH key for auth to the remote repository
  -login string
    	Login for HTTP auth to the remote repository
  -pass string
    	Password for HTTP auth to the remote repository
  -repo string
    	URL to remote repository
  -tag string
    	Remote tag for pull. NOTE (!): If the tag flag/env is specified, the specified branch flag/env will be ignored
  -timer string
    	Timeout for check update, default — 1m
```

* git-sync defaults to using environment variables if the flags are not explicitly set at startup:

| **env variable**   | **flag** | **example** |
|:---------------|:------|:--------|
|`GIT_SYNC_REPO` | -repo | `SSH - "git@github.com:eikoshelev/git-sync.git"` |
|                |       | `HTTP - "https://github.com/eikoshelev/git-sync.git"` |
|`GIT_SYNC_ROOT` | -dir | `"/path/to/your/folder/for/repo"` |
|`GIT_SYNC_WAIT` | -timer | `"1s", "2m", "3h", etc` |
|`GIT_SSH_KEY_PATH` | -key | `"/$HOME/.ssh/id_rsa"` |
|`SSH_KNOWN_HOSTS` | - | `"/$HOME/.ssh/known_hosts"`
|`GIT_HTTP_LOGIN` | -login | `<login for http>`
|`GIT_HTTP_PASSWORD` | -pass | `<password for http`
|`GIT_FORCE_PULL`  |  -force | `allowed - "true", not allowed - "false"` |
|`GIT_SYNC_BRANCH` | -branch | `"develop", "patch", etc`
|`GIT_SYNC_TAG` | -tag | `if the tag flag/env is specified, the branch flag/env will be ignored!`

### Docker container
```
docker pull eikoshelev/git-sync
```
