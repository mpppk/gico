package finder

import (
	"context"
	"github.com/mpppk/gico/gico"
	"github.com/google/go-github/github"
)

func SelectIssueInteractive(ctx context.Context, token string, remote *gico.Remote) (*github.Issue, error) {

	client := gico.GetGitHubClient(ctx, token)
	issues, err := gico.GetIssues(ctx, client, remote.Owner, remote.RepoName, nil)
	if err != nil {
		return nil, err
	}


	selectedIssueTitle, err := gico.PipeToPeco(gico.CreateIssueInfos(issues))
	if err != nil {
		return nil, err
	}

	if selectedIssueTitle == "" {
		return nil, nil
	}

	issue, err := gico.FindIssue(issues, selectedIssueTitle)
	if err != nil {
		return nil, err
	}

	return issue, nil
}
