[![Build Status](https://travis-ci.org/eikoshelev/git-sync.svg?branch=master)](https://travis-ci.org/eikoshelev/git-sync)
![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/eikoshelev/git-sync)

# git-sync

:recycle: Easy to use and stable update of your repository.
  
![alt text](assets/git-sync.png)
  
### Description
  
Real-time local repository update (every N time):
* **`Auth`**: authorization by HTTP (login/password) or SSH (private key);
* **`Clone`**: when starting, it checks the presence of the repository at the specified path (if the repository already exists, it starts checking for updates, if not, it will first clone);
* **`Fetch`**: checks repository updates every N time;
* **`Pull`**: as soon as the remote repository is updated - makes pull in the specified directory.

### Build and run locally
```sh
$ git clone https://github.com/eikoshelev/git-sync.git
$ cd git-sync && go build
```
### Docker container

The latest release is automatically published to the [Docker registry](https://hub.docker.com/r/eikoshelev/git-sync).

You can run it like this:
```sh
$ docker run -d --name git-sync eikoshelev/git-sync
```

### Usage
  
| **Environment variable** | **Description** |
| ------- | --- |
|`GSYNC_REPO_URL`        | URL to remote repository, example:  
`git@github.com:eikoshelev/git-sync.git` - for SSH  
`https://github.com/eikoshelev/git-sync.git` - for HTTP 
|`GSYNC_REPO_LOCAL_PATH` | Path to local directory for repository: `/path/to/your/local/directory` 
|`GSYNC_FETCH_TIMEOUT`   | Timeout for check update (fetch), default - `5m`, for example `1s`, `2m`, `3h`, etc 
|`GSYNC_SSH_KEY_PATH`    | Path to private SSH key for auth to the remote repository, for example - `/$HOME/.ssh/id_rsa` 
|`GSYNC_SSH_KNOWN_HOSTS` | Path to `known_hosts` file for work with remote repository, default - `/$HOME/.ssh/known_hosts`
|`GSYNC_HTTP_LOGIN`      | Login for HTTP auth to the remote repository 
|`GSYNC_HTTP_PASSWORD`   | Password for HTTP auth to the remote repository 
|`GSYNC_FORCE_PULL`      | Forced pull with changed local repository: allowed - `true`, not allowed (default) - `false` 
|`GSYNC_REPO_BRANCH`     | Remote branch for pull, for example - `develop`, `patch`, etc.  
**NOTE: If the 'tag' is specified, the 'branch' will be ignored!**
|`GSYNC_REPO_TAG`        | Remote tag for clone, for example - `"v1.0.0"`, `"v2.0-rc"`, `"v3.0-stable"`, etc.  
**NOTE: If the 'tag' is specified, the specified 'branch' will be ignored!**

### What else?

Open an issue or PR if you have more suggestions, questions or ideas about what to add.
