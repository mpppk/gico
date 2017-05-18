package finder

import (
	"github.com/mpppk/hlb/project"
	"strconv"
)

type FilterableIssue struct {
	project.Issue
}


func (f *FilterableIssue) FilterString() string {
	return "#" + strconv.Itoa(f.GetNumber()) + " " + f.GetTitle()
}
