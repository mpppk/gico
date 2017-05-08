package project

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"errors"
	"net/url"
)

type GitHubService struct {
	Client *github.Client
}

func NewGitHubService(ctx context.Context, token string, baseUrlStrs ...string) (*GitHubService, error) {
	if len(baseUrlStrs) > 1 {
		return nil, errors.New("too many base urls")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	if len(baseUrlStrs) == 1 {
		baseUrl, err := url.Parse(baseUrlStrs[0])
		if err != nil {
			return nil, err
		}

		client.BaseURL = baseUrl
	}

	return &GitHubService{Client: client}, nil
}

func (s *GitHubService) GetPullRequests(ctx context.Context, owner, repo string) (pullRequests []PullRequest, err error) {
	gitHubPullRequests, _, err := s.Client.PullRequests.List(ctx, owner, repo, nil)

	if err != nil {
		return nil, err
	}

	for _, gitHubPullRequest := range gitHubPullRequests {
		pullRequests = append(pullRequests, Issue(&GitHubPullRequest{PullRequest: gitHubPullRequest}))
	}

	return pullRequests, err
}

func (s *GitHubService) GetIssues(ctx context.Context, owner, repo string) (issues []Issue, err error) {
	gitHubIssues, err := s.getGitHubIssues(ctx, s.Client, owner, repo, nil)

	if err != nil {
		return nil, err
	}

	for _, gitHubIssue := range gitHubIssues {
		issues = append(issues, Issue(&GitHubIssue{Issue: gitHubIssue}))
	}

	return issues, err
}

func (s *GitHubService) getGitHubIssues(ctx context.Context, client *github.Client, owner, repo string, opt *github.IssueListByRepoOptions) (issues []*github.Issue, err error) {
	issuesAndPRs, _, err := client.Issues.ListByRepo(ctx, owner, repo, opt)

	if err != nil {
		return nil, err
	}

	for _, issueOrPR := range issuesAndPRs {
		if issueOrPR.PullRequestLinks == nil {
			issues = append(issues, issueOrPR)
		}
	}
	return issues, nil
}

func (s *GitHubService) GetRepository(ctx context.Context, owner, repo string) (Repository, error) {
	githubRepo, _, err := s.Client.Repositories.Get(ctx, owner, repo)

	if err != nil {
		return nil, err
	}

	return Repository(&GitHubRepository{Repository: githubRepo}), err
}

