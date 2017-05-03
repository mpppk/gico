package project

import (
	"github.com/xanzy/go-gitlab"
	"context"
	"errors"
	githubw "github.com/mpppk/gico/github"
	"github.com/mpppk/gico/etc"
)

func GetRepository(ctx context.Context, hostType, token, owner, repo string) (Repository, error) {
	switch hostType {
	case etc.HOST_TYPE_GITHUB.String():
		client := githubw.GetGitHubClient(ctx, token)
		repo, _, err := client.Repositories.Get(ctx, owner, repo)

		if err != nil {
			return nil, err
		}

		return Repository(&GitHubRepository{Repository: repo}), err

	case etc.HOST_TYPE_GITLAB.String():
		client := gitlab.NewClient(nil, token)
		gitLabProject, _, err := client.Projects.GetProject(owner + "/" + repo)

		if err != nil {
			return nil, err
		}

		return Repository(&GitLabRepository{Project: gitLabProject}), err
	}

	return nil, errors.New("unknown host type")
}