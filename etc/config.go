package etc

type HostType int

const (
	HOST_TYPE_GITHUB HostType = iota + 1
	HOST_TYPE_GITLAB
	HOST_TYPE_BITBUCKET
	HOST_TYPE_GITBUCKET
)

func (s HostType) String() string {
	switch s {
	case HOST_TYPE_GITHUB:
		return "github"
	case HOST_TYPE_GITLAB:
		return "gitlab"
	case HOST_TYPE_BITBUCKET:
		return "bitbucket"
	case HOST_TYPE_GITBUCKET:
		return "gitbucket"
	default:
		return "Unknown"
	}
}

type Host struct {
	Host       string
	HostType   string `mapstructure:"host_type"`
	OAuthToken string `mapstructure:"oauth_token"`
}

type Config struct {
	Hosts []*Host
}
