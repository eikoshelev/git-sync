package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// git
	SSHKeyPath   string `envconfig:"GSYNC_SSH_KEY_PATH"`
	HTTPLogin    string `envconfig:"GSYNC_HTTP_LOGIN"`
	HTTPPassword string `envconfig:"GSYNC_HTTP_PASSWORD"`
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
	return &c, err
}
