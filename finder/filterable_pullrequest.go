package finder

import (
	"github.com/mpppk/hlb/project"
	"strconv"
)

type FilterablePullRequest struct {
	project.PullRequest
}

func (f *FilterablePullRequest) FilterString() string {
	return "!" + strconv.Itoa(f.GetNumber()) + " " + f.GetTitle()
}

func toFilterableFromPullRequests(prs []*FilterablePullRequest) (strs []FilterableStringer) {
	for _, pr := range prs {
		strs = append(strs, FilterableStringer(pr))
	}
	return strs
}
