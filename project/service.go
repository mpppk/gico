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

func GetService(ctx context.Context, hostType, token string) (Service, error){
	switch hostType {
	case etc.HOST_TYPE_GITHUB.String():
		return Service(NewGitHubService(ctx, token)), nil
	}
	switch hostType {
	case etc.HOST_TYPE_GITLAB.String():
		return Service(NewGitLabService(token)), nil
	}
	return nil, errors.New("unknown host type: " + hostType)
}
