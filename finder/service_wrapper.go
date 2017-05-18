package finder

import (
	"github.com/mpppk/hlb/hlb"
)

type ServiceWrapper struct {
	*hlb.ServiceWrapper
}

func (s *ServiceWrapper) GetFilterableIssues() ([]*FilterableIssue, error) {
	issues, err := s.GetIssues()
	if err != nil {
		return nil, err
	}

	var filterableIssues []*FilterableIssue
	for _, is := range issues {
		filterableIssues = append(filterableIssues, &FilterableIssue{Issue: is})
	}
	return filterableIssues, nil
}

func (s *ServiceWrapper) GetFilterablePullRequests() ([]*FilterablePullRequest, error) {
	prs, err := s.GetPullRequests()
	if err != nil {
		return nil, err
	}

	var filterablePrs []*FilterablePullRequest
	for _, is := range prs {
		filterablePrs = append(filterablePrs, &FilterablePullRequest{PullRequest: is})
	}
	return filterablePrs, nil
}

func NewServiceWrapper(base *hlb.CmdBase) *ServiceWrapper {
	return &ServiceWrapper{&hlb.ServiceWrapper{Base: base}}
}
