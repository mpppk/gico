package project

import (
	"github.com/xanzy/go-gitlab"
	"github.com/google/go-github/github"
	"strconv"
	"errors"
)

type PullRequest interface {
	GetNumber() int
	GetTitle() string
	GetHTMLURL() string
}

type GitHubPullRequest struct {
	*github.PullRequest
}

type GitLabPullRequest struct {
	*gitlab.MergeRequest
}

func (pullRequest *GitLabPullRequest) GetNumber() int {
	return pullRequest.IID
}

func (pullRequest *GitLabPullRequest) GetTitle() string {
	return pullRequest.Title
}

func (pullRequest *GitLabPullRequest) GetHTMLURL() string {
	return pullRequest.WebURL
}

func createPullRequestInfo(pullRequest PullRequest) string {
	return "#" + strconv.Itoa(pullRequest.GetNumber()) + " " + pullRequest.GetTitle()
}

func CreatePullRequestInfos(pullRequests []PullRequest) (pullRequestInfos []string) {
	for _, pullRequest := range pullRequests {
		pullRequestInfos = append(pullRequestInfos, createPullRequestInfo(pullRequest))
	}
	return pullRequestInfos
}

func FindPullRequest(pullRequests []PullRequest, pullRequestInfo string) (PullRequest, error) {
	var targetPullRequest PullRequest = nil
	for _, pullRequest := range pullRequests {
		if createPullRequestInfo(pullRequest) == pullRequestInfo {
			targetPullRequest = pullRequest
			break
		}
	}

	if targetPullRequest == nil {
		return nil, errors.New("pullRequest not found")
	}

	return targetPullRequest, nil
}

