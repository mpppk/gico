package finder

import (
	"context"
	"github.com/mpppk/gico/git"
	"github.com/mpppk/gico/project"
	"github.com/mpppk/gico/utils"
)

func SelectIssueInteractive(ctx context.Context, service project.Service, remote *git.Remote) (project.Issue, error) {

	issues, err := service.GetIssues(ctx, remote.Owner, remote.RepoName)

	if err != nil {
		return nil, err
	}

	selectedIssueTitle, err := utils.PipeToPeco(project.CreateIssueInfos(issues))
	if err != nil {
		return nil, err
	}

	if selectedIssueTitle == "" {
		return nil, nil
	}

	return project.FindIssue(issues, selectedIssueTitle)
}

func SelectPullRequestInteractive(ctx context.Context, service project.Service, remote *git.Remote) (project.Issue, error) {
	prs, err := service.GetPullRequests(ctx, remote.Owner, remote.RepoName)

	if err != nil {
		return nil, err
	}

	selectedPullRequestTitle, err := utils.PipeToPeco(project.CreatePullRequestInfos(prs))
	if err != nil {
		return nil, err
	}

	if selectedPullRequestTitle == "" {
		return nil, nil
	}

	return project.FindPullRequest(prs, selectedPullRequestTitle)
}