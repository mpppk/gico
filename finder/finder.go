package finder

import (
	"github.com/mpppk/gico/utils"
	"errors"
)

type FilterableStringer interface {
	FilterString() string
}

func Filter(fss []FilterableStringer) (FilterableStringer, error) {
	var filterabletStrs []string
	for _, fs := range fss {
		filterabletStrs = append(filterabletStrs, fs.FilterString())
	}
	info, err := utils.PipeToPeco(filterabletStrs)
	if err != nil {
		return nil, err
	}
	for _, fs := range fss {
		if fs.FilterString() == info {
			return fs, nil
		}
	}
	return nil, errors.New("not found")
}

func FilterIssues(issues []*FilterableIssue) (*FilterableIssue, error) {
	var filterableStrings []FilterableStringer
	for _, is := range issues {
		filterableStrings = append(filterableStrings, FilterableStringer(is))
	}

	res, err := Filter(filterableStrings)
	if err != nil {
		return nil, err
	}
	return res.(*FilterableIssue), nil
}

func FilterPullRequests(prs []*FilterablePullRequest) (*FilterablePullRequest, error) {
	var filterableStrings []FilterableStringer
	for _, pr := range prs {
		filterableStrings = append(filterableStrings, FilterableStringer(pr))
	}

	res, err := Filter(filterableStrings)
	if err != nil {
		return nil, err
	}
	return res.(*FilterablePullRequest), nil
}