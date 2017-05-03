package finder

import (
	"github.com/mpppk/gico/git"
	"github.com/xanzy/go-gitlab"
	"github.com/mpppk/gico/utils"
	"errors"
)

func SelectGitLabIssueInteractive(token string, remote *git.Remote) (*gitlab.Issue, error) {

	client := gitlab.NewClient(nil, token)
	issues, _, err := client.Issues.ListProjectIssues("mpppk/test", nil)
	if err != nil {
		return nil, err
	}

	var issueTitles []string
	for _, issue := range issues {
		issueTitles = append(issueTitles, issue.Title)
	}

	selectedIssueTitle, err := utils.PipeToPeco(issueTitles)
	if err != nil {
		return nil, err
	}

	if selectedIssueTitle == "" {
		return nil, nil
	}

	for _, issue := range issues {
		if issue.Title == selectedIssueTitle {
			return issue, nil
		}
	}
	return nil, errors.New("issue not fuond")
}
