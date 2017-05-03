package finder

import (
	"context"
	"github.com/mpppk/gico/git"
	"github.com/mpppk/gico/project"
	"github.com/mpppk/gico/utils"
)

func SelectIssueInteractive(ctx context.Context, hostType, token string, remote *git.Remote) (project.Issue, error) {
	issues, err := project.GetIssues(ctx, hostType, token, remote.Owner, remote.RepoName)

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