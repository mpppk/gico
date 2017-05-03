package project

import (
	"github.com/google/go-github/github"
	"github.com/xanzy/go-gitlab"
)

type Repository interface {
	GetHTMLURL() string
}

type GitHubRepository struct {
	*github.Repository
}

type GitLabRepository struct {
	*gitlab.Project
}

func (repo *GitLabRepository) GetHTMLURL() string {
	return repo.WebURL
}
