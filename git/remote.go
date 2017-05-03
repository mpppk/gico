package git

import (
	"os/exec"
	"strings"
	"errors"
	"github.com/mpppk/gico/utils"
	"github.com/mpppk/gico/etc"
)

type Remote struct {
	Name     string
	FetchURL string
	PushURL  string
	HostType string
	Owner    string
	RepoName string
}

func GetOriginRemote(hosts []*etc.Host) (*Remote, error) {
	out, err := exec.Command("git", "remote", "-v").Output()

	if err != nil {
		return nil, err
	}

	remoteStrs := strings.Split(string(out), "\n")
	remoteStrs = remoteStrs[:len(remoteStrs)-1]

	originRemoteStrs, ok := findOriginRemoteStrs(remoteStrs)
	if !ok {
		return nil, errors.New("origin remote not found")

	}

	var originRemote = &Remote{}
	for _, remoteStr := range originRemoteStrs {
		remoteNameAndUrlInfo := strings.Split(remoteStr, "\t")

		originRemote.Name = remoteNameAndUrlInfo[0]
		urlInfo := strings.Split(remoteNameAndUrlInfo[1], " ")

		host, owner, repoName, err := utils.ParseRemoteURL(urlInfo[0], hosts)

		if err != nil {
			return nil, err
		}

		originRemote.HostType = host
		originRemote.Owner = owner
		originRemote.RepoName = repoName

		if urlInfo[1] == "(fetch)" {
			originRemote.FetchURL = urlInfo[0]
		} else if urlInfo[1] == "(push)" {
			originRemote.PushURL = urlInfo[0]
		} else {
			return nil, errors.New("unknown origin remote url found")
		}
	}
	return originRemote, nil
}

func findOriginRemoteStrs(remotes []string) (remoteStrs []string, ok bool) {
	for _, remoteStr := range remotes {
		if strings.Contains(remoteStr, "origin") {
			remoteStrs = append(remoteStrs, remoteStr)
		}
	}

	return remoteStrs, len(remoteStrs) > 0
}
