package finder

import (
	"github.com/mpppk/gico/utils"
	"errors"
)

type FilterableStringer interface {
	FilterString() string
}

type FilterStringer interface {
	FilterString() string
}

type FilterableLink interface {
	FilterStringer
	GetHTMLURL() string
}

type FilterableUrl struct {
	Url string
	FilterStr string
}

func (f *FilterableUrl) FilterString() string{
	return f.FilterStr
}

func (f *FilterableUrl) GetHTMLURL() string{
	return f.Url
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
	res, err := Filter(toFilterableFromIssues(issues))
	return res.(*FilterableIssue), err
}

func FilterPullRequests(prs []*FilterablePullRequest) (*FilterablePullRequest, error) {
	res, err := Filter(toFilterableFromPullRequests(prs))
	return res.(*FilterablePullRequest), err
}

func FilterLinks(links []FilterableLink) (FilterableLink, error) {
	var strs []FilterableStringer
	for _, link := range links {
		strs = append(strs, FilterableStringer(link))
	}

	res, err := Filter(strs)
	return res.(FilterableLink), err
}