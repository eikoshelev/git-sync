[![Build Status](https://travis-ci.org/eikoshelev/git-sync.svg?branch=master)](https://travis-ci.org/eikoshelev/git-sync)
![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/eikoshelev/git-sync)

# git-sync

:recycle: Easy to use and stable update of your repository.
  
![alt text](assets/git-sync.png)
  
### Description
  
Real-time local repository update (every N time):
* Authorization by HTTP (login/password) or SSH (private key);
* When starting, it checks the presence of the repository at the specified path (if the repository already exists, it starts checking for updates, if not, it will first clone);
* Checks repository updates every N time;
* As soon as the remote repository is updated - makes ```git pull``` in the specified directory.

### Build and run locally
```sh
$ git clone https://github.com/eikoshelev/git-sync.git
$ cd git-sync && go build
$ ./git-sync <flags>
```
### Docker container

The latest release is automatically published to the [Docker registry](https://hub.docker.com/r/eikoshelev/git-sync).

You can run it like this:
```sh
$ docker run -d --name git-sync eikoshelev/git-sync
```

### Usage

* Description of the flags used: ```./git-sync -h```

* **git-sync** defaults to using environment variables if the flags are not explicitly set at startup:
  
| **Environment Variable** | **Flag** | **Description** |
| --- | --- | --- |
|`GIT_SYNC_REPO`    | -repo   | URL to remote repository, example: `git@github.com:eikoshelev/git-sync.git` - for SSH, `https://github.com/eikoshelev/git-sync.git` - for HTTP 
|`GIT_SYNC_ROOT`    | -dir    | Path to local directory for repository: `/path/to/your/local/directory` 
|`GIT_SYNC_WAIT`    | -timer  | Timeout for check update (fetch), default — `1m`, for example `1s`, `2m`, `3h`, etc 
|`GIT_SSH_KEY_PATH` | -key    | Path to private SSH key for auth to the remote repository, for example - `/$HOME/.ssh/id_rsa` 
|`SSH_KNOWN_HOSTS`  | -       | Path to `known_hosts` file for work with remote repository, default - `/$HOME/.ssh/known_hosts`
|`GIT_HTTP_LOGIN`   | -login  | Login for HTTP auth to the remote repository 
|`GIT_HTTP_PASSWORD`| -pass   | Password for HTTP auth to the remote repository 
|`GIT_FORCE_PULL`   | -force  | Forced pull with changed local repository: allowed - `true`, not allowed (default) - `false` 
|`GIT_SYNC_BRANCH`  | -branch | Remote branch for pull, for example - `develop`, `patch`, etc. **NOTE: If the 'tag' flag/env is specified, the 'branch' flag/env will be ignored!**
|`GIT_SYNC_TAG`     | -tag    | Remote tag for pull, for example - `"v1.0.0"`, `"v2.0"`, `"v3.0-stable"`, etc. **NOTE: If the tag flag/env is specified, the specified branch flag/env will be ignored!**

### What else?

Open an issue or PR if you have more suggestions, questions or ideas about what to add.
