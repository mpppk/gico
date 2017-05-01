package finder

import (
	"github.com/mpppk/gico/gico"
	"strings"
	"fmt"
)

func GetBranchInteractive() (string, error) {
	names, err := gico.GetBranchNames()

	str, err := gico.PipeToPeco(names)

	if err != nil {
		return "", err
	}

	if len(str) == 0 {
		return "", nil
	}

	branchName := strings.Trim(string(str), " \n")
	return gico.TrimBranchName(branchName), nil
}

func SwitchBranchInteractive() error {
	branchName, err := GetBranchInteractive()

	if err != nil {
		fmt.Println(err)
	}

	err = gico.ExecCommand("git", "checkout", branchName)

	if err != nil {
		return err
	}
	return nil
}

func GetLogHashInteractive() (string, error) {
	logs, err := gico.GetLogs()

	if err != nil {
		return "", err
	}

	logs, err = gico.ArrangeHashPosition(logs)

	if err != nil {
		return "", err
	}

	log, err := gico.PipeToPeco(logs)

	if err != nil {
		return "", err
	}

	return gico.ExtractHash(log)
}
