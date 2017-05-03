package project

import (
	"github.com/xanzy/go-gitlab"
	"strconv"
	"errors"
	"github.com/google/go-github/github"
)

type Issue interface {
	GetNumber() int
	GetTitle() string
	GetHTMLURL() string
}

type GitHubIssue struct {
	*github.Issue
}

type GitLabIssue struct {
	*gitlab.Issue
}

func (issue *GitLabIssue) GetNumber() int {
	return issue.IID
}

func (issue *GitLabIssue) GetTitle() string {
	return issue.Title
}

func (issue *GitLabIssue) GetHTMLURL() string {
	return issue.WebURL
}

func createIssueInfo(issue Issue) string {
	return "#" + strconv.Itoa(issue.GetNumber()) + " " + issue.GetTitle()
}

func CreateIssueInfos(issues []Issue) (issueInfos []string) {
	for _, issue := range issues {
		issueInfos = append(issueInfos, createIssueInfo(issue))
	}
	return issueInfos
}

func FindIssue(issues []Issue, issueInfo string) (Issue, error) {
	var targetIssue Issue = nil
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

