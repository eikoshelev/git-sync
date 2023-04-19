package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	HomeDirectory string `envconfig:"HOME"`
	// git
	SSHKeyPath    string `envconfig:"GSYNC_SSH_KEY_PATH"`    // default: "/$HOME/.ssh/id_rsa"
	SSHKnownHosts string `envconfig:"GSYNC_SSH_KNOWN_HOSTS"` // default: "/$HOME/.ssh/known_hosts"
	HTTPLogin     string `envconfig:"GSYNC_HTTP_LOGIN"`
	HTTPPassword  string `envconfig:"GSYNC_HTTP_PASSWORD"`
	// repository
	RemoteRepoURL    string `envconfig:"GSYNC_REPO_URL" required:"true"`
	RemoteRepoTag    string `envconfig:"GSYNC_REPO_TAG"`
	RemoteRepoBranch string `envconfig:"GSYNC_REPO_BRANCH" required:"true" default:"master"`
	LocalRepoPath    string `envconfig:"GSYNC_REPO_LOCAL_PATH" required:"true"`
	// options
	FetchTimeout time.Duration `envconfig:"GSYNC_FETCH_TIMEOUT" default:"5m"`
	ForcePull    bool          `envconfig:"GSYNC_FORCE_PULL" default:"false"`
}

func GetConfiguration() (*config, error) {
	var c config
	err := envconfig.Process("GSYNC", &c)
	if err != nil {
		return nil, err
	}
	c.Validate()
	return &c, nil
}

const (
	// repo tag/branch
	tagReference    = "refs/tags/%s"
	branchReference = "refs/heads/%s"
	defaultBranch   = "master"
	// local file paths
	sshKeyPathTempl    = "%s/.ssh/id_rsa"
	sshKnownHostsTempl = "%s/.ssh/known_hosts"
)

func (c *config) Validate() {
	if c.RemoteRepoTag != "" {
		c.RemoteRepoBranch = fmt.Sprintf(tagReference, c.RemoteRepoTag)
	} else {
		if c.RemoteRepoBranch != "" {
			c.RemoteRepoBranch = fmt.Sprintf(branchReference, c.RemoteRepoBranch)
		} else {
			c.RemoteRepoBranch = fmt.Sprintf(branchReference, defaultBranch)
		}
	}
	if c.SSHKeyPath == "" {
		c.SSHKeyPath = fmt.Sprintf(sshKeyPathTempl, c.HomeDirectory)
	}
	if c.SSHKnownHosts == "" {
		c.SSHKnownHosts = fmt.Sprintf(sshKnownHostsTempl, c.HomeDirectory)
	}
}
