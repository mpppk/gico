package gico

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"github.com/skratchdot/open-golang/open"
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

func FindIssue(issues []*github.Issue, title string) (*github.Issue, error) {
	var targetIssue *github.Issue = nil
	for _, issue := range issues {
		if issue.GetTitle() == title {
			targetIssue = issue
			break
		}
	}

	if targetIssue == nil {
		return nil, errors.New("issue not found")
	}

	return targetIssue, nil
}

func OpenIssuePageInteractive(ctx context.Context, token string, remote *Remote) {
	client := GetGitHubClient(ctx, token)
	issues, err := GetIssues(ctx, client, remote.Owner, remote.RepoName, nil)
	PanicIfErrorExist(err)

	var issueTitles []string
	for _, issue := range issues {
		issueTitles = append(issueTitles, issue.GetTitle())
	}

	selectedIssueTitle, err := PipeToPeco(issueTitles)
	PanicIfErrorExist(err)

	issue, err := FindIssue(issues, selectedIssueTitle)
	PanicIfErrorExist(err)

	open.Run(issue.GetHTMLURL())
}