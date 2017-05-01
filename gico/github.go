package gico

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"strconv"
)

func GetGitHubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func GetIssues(ctx context.Context, client *github.Client, owner, repo string, opt *github.IssueListByRepoOptions) ([]*github.Issue, error) {
	issuesAndPRs, _, err := client.Issues.ListByRepo(ctx, owner, repo, opt)

	if err != nil {
		return nil, err
	}

	var issues []*github.Issue
	for _, issueOrPR := range issuesAndPRs {
		if issueOrPR.PullRequestLinks == nil {
			issues = append(issues, issueOrPR)
		}
	}
	return issues, nil
}

func createIssueInfo(issue *github.Issue) string {
	return "#" + strconv.Itoa(issue.GetNumber()) + " " + issue.GetTitle()
}

func CreateIssueInfos(issues []*github.Issue) (issueInfos []string) {
	for _, issue := range issues {
		issueInfos = append(issueInfos, createIssueInfo(issue))
	}
	return issueInfos
}

func FindIssue(issues []*github.Issue, issueInfo string) (*github.Issue, error) {
	var targetIssue *github.Issue = nil
	for _, issue := range issues {
		if createIssueInfo(issue) == issueInfo {
			targetIssue = issue
			break
		}
	}

	if targetIssue == nil {
		return nil, errors.New("issue not found")
	}

	return targetIssue, nil
}

