package finder

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/mpppk/gico/git"
	ggithub "github.com/mpppk/gico/github"
	"github.com/mpppk/gico/utils"
)

func SelectIssueInteractive(ctx context.Context, token string, remote *git.Remote) (*github.Issue, error) {

	client := ggithub.GetGitHubClient(ctx, token)
	issues, err := ggithub.GetIssues(ctx, client, remote.Owner, remote.RepoName, nil)
	if err != nil {
		return nil, err
	}

	selectedIssueTitle, err := utils.PipeToPeco(ggithub.CreateIssueInfos(issues))
	if err != nil {
		return nil, err
	}

	if selectedIssueTitle == "" {
		return nil, nil
	}

	issue, err := ggithub.FindIssue(issues, selectedIssueTitle)
	if err != nil {
		return nil, err
	}

	return issue, nil
}
