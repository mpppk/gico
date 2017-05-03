package project

import (
	"context"
	githubw "github.com/mpppk/gico/github"
	"github.com/mpppk/gico/etc"
	"github.com/xanzy/go-gitlab"
	"errors"
)

func GetPullRequests(ctx context.Context, hostType, token, owner, repo string) (pullRequests []PullRequest, err error) {
	switch hostType {
	case etc.HOST_TYPE_GITHUB.String():
		client := githubw.GetGitHubClient(ctx, token)
		gitHubPullRequests, _, err := client.PullRequests.List(ctx, owner, repo, nil)

		if err != nil {
			return nil, err
		}

		for _, gitHubPullRequest := range gitHubPullRequests {
			pullRequests = append(pullRequests, Issue(&GitHubPullRequest{PullRequest: gitHubPullRequest}))
		}

		return pullRequests, err

	case etc.HOST_TYPE_GITLAB.String():
		client := gitlab.NewClient(nil, token)
		gitLabMergeRequests, _, err := client.MergeRequests.ListMergeRequests(owner + "/" + repo, nil)

		for _, gitLabMergeRequest := range gitLabMergeRequests {
			pullRequests = append(pullRequests, Issue(&GitLabPullRequest{MergeRequest:gitLabMergeRequest}))
		}

		return pullRequests, err
	}

	return nil, errors.New("unknown host type")
}
