package main

import (
	"fmt"
	"gopkg.in/libgit2/git2go.v25"
)

func main() {
	repoPath := "../ae-web-preview"
	r, err := git.OpenRepository(repoPath)
	if err != nil {
		fmt.Println(err)
	}

	branchIterator, err := r.NewBranchIterator(git.BranchLocal)

	var branchIteratorFunc git.BranchIteratorFunc = func(branch *git.Branch, branchType git.BranchType) error {
		name, err := branch.Name()
		if err != nil {
			return err
		}
		fmt.Println(name)
		return nil
	}

	branchIterator.ForEach(branchIteratorFunc)
}
