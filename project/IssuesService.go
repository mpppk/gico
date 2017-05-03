package project

import (
	"context"
	githubw "github.com/mpppk/gico/github"
	"github.com/mpppk/gico/etc"
	"github.com/xanzy/go-gitlab"
	"errors"
)

func GetIssues(ctx context.Context, hostType, token, owner, repo string) (issues []Issue, err error) {
	switch hostType {
	case etc.HOST_TYPE_GITHUB.String():
		client := githubw.GetGitHubClient(ctx, token)
		gitHubIssues, err := githubw.GetIssues(ctx, client, owner, repo, nil)

		if err != nil {
			return nil, err
		}

		for _, gitHubIssue := range gitHubIssues {
			issues = append(issues, Issue(&GitHubIssue{Issue: gitHubIssue}))
		}

		return issues, err

	case etc.HOST_TYPE_GITLAB.String():
		client := gitlab.NewClient(nil, token)
		gitLabIssues, _, err := client.Issues.ListProjectIssues(owner + "/" + repo, nil)

		for _, gitLabIssue := range gitLabIssues {
			issues = append(issues, Issue(&GitLabIssue{Issue:gitLabIssue}))
		}

		return issues, err
	}

	return nil, errors.New("unknown host type")
}

