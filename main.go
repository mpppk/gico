package main

import (
	"fmt"
	"gopkg.in/libgit2/git2go.v25"
	"github.com/mattn/go-pipeline"
	"strings"
	"os/exec"
	"os"
)

func main() {
	repoPath := "."
	r, err := git.OpenRepository(repoPath)
	if err != nil {
		fmt.Println(err)
	}

	err = switchBranch(r)

	if err != nil {
		fmt.Println(err)
	}
}

func switchBranch(repo *git.Repository) error {
	names, err := getBranchNames(repo)

	out, err := pipeline.Output(
		[]string{"echo", strings.Join(names, "\n")},
		[]string{"peco"},
	)

	if err != nil {
		return err
	}

	branchName := strings.Trim(string(out), " \n")
	fmt.Println("branch ", branchName)
	cmd := exec.Command("git", "checkout", branchName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func getBranchNames(repo *git.Repository) ([]string, error) {
	branchIterator, err := repo.NewBranchIterator(git.BranchLocal)
	if err != nil {
		return nil, err
	}

	var branchNames []string
	var branchIteratorFunc git.BranchIteratorFunc = func(branch *git.Branch, branchType git.BranchType) error {
		name, err := branch.Name()
		if err != nil {
			return err
		}
		branchNames = append(branchNames, name)
		return nil
	}

	branchIterator.ForEach(branchIteratorFunc)
	return branchNames, nil
}
