package project

import (
	"context"
	"github.com/xanzy/go-gitlab"
)

type GitLabService struct {
	Client *gitlab.Client
}

func NewGitLabService(token string) *GitLabService {
	return &GitLabService{Client: gitlab.NewClient(nil, token)}
}

func (s *GitLabService) GetIssues(ctx context.Context, owner, repo string) (issues []Issue, err error) {
	gitLabIssues, _, err := s.Client.Issues.ListProjectIssues(owner + "/" + repo, nil)

	for _, gitLabIssue := range gitLabIssues {
		issues = append(issues, Issue(&GitLabIssue{Issue:gitLabIssue}))
	}

	return issues, err
}

func (s *GitLabService) GetPullRequests(ctx context.Context, owner, repo string) (pullRequests []PullRequest, err error) {
	gitLabMergeRequests, _, err := s.Client.MergeRequests.ListMergeRequests(owner + "/" + repo, nil)

	for _, gitLabMergeRequest := range gitLabMergeRequests {
		pullRequests = append(pullRequests, Issue(&GitLabPullRequest{MergeRequest:gitLabMergeRequest}))
	}

	return pullRequests, err
}

func (s *GitLabService) GetRepository(ctx context.Context, owner, repo string) (Repository, error) {
	gitLabProject, _, err := s.Client.Projects.GetProject(owner + "/" + repo)

	if err != nil {
		return nil, err
	}

	return Repository(&GitLabRepository{Project: gitLabProject}), err
}
