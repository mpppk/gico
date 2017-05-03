package github

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetGitHubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func GetIssues(ctx context.Context, client *github.Client, owner, repo string, opt *github.IssueListByRepoOptions) (issues []*github.Issue, err error) {
	issuesAndPRs, _, err := client.Issues.ListByRepo(ctx, owner, repo, opt)

	if err != nil {
		return nil, err
	}

	for _, issueOrPR := range issuesAndPRs {
		if issueOrPR.PullRequestLinks == nil {
			//ghIssue := project.GitHubIssue{issueOrPR}
			issues = append(issues, issueOrPR)
		}
	}
	return issues, nil
}
