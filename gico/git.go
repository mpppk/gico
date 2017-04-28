package gico

import (
	"fmt"
	"gopkg.in/libgit2/git2go.v25"
	"github.com/mattn/go-pipeline"
	"strings"
	"os/exec"
	"os"
)

func SwitchBranch(repo *git.Repository) error {
	names, err := getBranchNames()

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

func getBranchNames() ([]string, error) {
	out, err := exec.Command("git", "branch").Output()

	if err != nil {
		return nil, err
	}

	branchNames := strings.Split(string(out), "\n")

	for i, b := range branchNames {
		branchNames[i] = strings.Trim(b, "* \n")
	}

	return branchNames, nil
}
