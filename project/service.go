package project

import (
	"context"
	"github.com/mpppk/gico/etc"
	"errors"
)

type Service interface{
	GetPullRequests(ctx context.Context, owner, repo string) ([]PullRequest, error)
	GetIssues(ctx context.Context, owner, repo string) ([]Issue, error)
	GetRepository(ctx context.Context, owner, repo string) (Repository, error)
}

func GetService(ctx context.Context, host, hostType, token string) (Service, error){
	switch hostType {
	case etc.HOST_TYPE_GITHUB.String():
		service, err := NewGitHubService(ctx, token, "https://api." + host)
		if err != nil {
			return nil, err
		}
		return Service(service), nil
	}
	switch hostType {
	case etc.HOST_TYPE_GITLAB.String():
		service, err := NewGitLabService(token, "https://" + host + "/api/v3")
		if err != nil {
			return nil, err
		}

		return Service(service), nil
	}
	return nil, errors.New("unknown host type: " + hostType)
}