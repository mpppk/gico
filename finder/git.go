package finder

import (
	"strings"
	"fmt"
	"github.com/mpppk/gico/git"
	"github.com/mpppk/gico/utils"
)

func GetBranchInteractive() (string, error) {
	names, err := git.GetBranchNames()

	str, err := utils.PipeToPeco(names)

	if err != nil {
		return "", err
	}

	if len(str) == 0 {
		return "", nil
	}

	branchName := strings.Trim(string(str), " \n")
	return git.TrimBranchName(branchName), nil
}

func SwitchBranchInteractive() error {
	branchName, err := GetBranchInteractive()

	if err != nil {
		fmt.Println(err)
	}

	err = utils.ExecCommand("git", "checkout", branchName)

	if err != nil {
		return err
	}
	return nil
}

func GetLogHashInteractive() (string, error) {
	logs, err := git.GetLogs()

	if err != nil {
		return "", err
	}

	logs, err = git.ArrangeHashPosition(logs)

	if err != nil {
		return "", err
	}

	log, err := utils.PipeToPeco(logs)

	if err != nil {
		return "", err
	}

	return git.ExtractHash(log)
}
