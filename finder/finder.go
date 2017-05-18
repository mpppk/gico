package finder

import (
	"github.com/mpppk/hlb/project"
	"strconv"
	"github.com/mpppk/gico/utils"
	"errors"
)

type FilterableStringer interface {
	FilterString() string
}

type FilterableIssue struct {
	project.Issue
}

type FilterablePullRequest struct {
	project.PullRequest
}

func (f *FilterableIssue) FilterString() string {
	return "#" + strconv.Itoa(f.GetNumber()) + " " + f.GetTitle()
}

func (f *FilterablePullRequest) FilterString() string {
	return "!" + strconv.Itoa(f.GetNumber()) + " " + f.GetTitle()
}

func Filter(fss []FilterableStringer) (FilterableStringer, error) {
	var infos []string
	for _, fs := range fss {
		infos = append(infos, fs.FilterString())
	}
	info, err := utils.PipeToPeco(infos)
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
