package gico

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func SwitchBranch() error {
	branchName, err := GetBranchInteractive()

	if err != nil {
		fmt.Println(err)
	}

	err = execCommand("git", "checkout", branchName)

	if err != nil {
		return err
	}
	return nil
}

func GetBranchInteractive() (string, error) {
	names, err := getBranchNames()

	str, err := PipeToPeco(names)

	if err != nil {
		return "", err
	}

	if len(str) == 0 {
		return "", nil
	}

	branchName := strings.Trim(string(str), " \n")
	return trimBranchName(branchName), nil
}

func getBranchNames() ([]string, error) {
	out, err := exec.Command("git", "branch", "-a").Output()

	if err != nil {
		return nil, err
	}

	branchNames := strings.Split(string(out), "\n")

	return branchNames, nil
}

func extractHash(log string) (string, error) {
	reg := regexp.MustCompile(`[0123456789abcdef]{7}`)
	return reg.FindString(log), nil
}

func arrangeHashPosition(logs []string) ([]string, error) {
	var newLogs []string
	for _, log := range logs {
		hash, err := extractHash(log)

		if err != nil {
			return nil, err
		}

		newLog := strings.Replace(log, hash, "", -1)

		newLogs = append(newLogs, newLog+" ["+hash+"]")
	}
	return newLogs, nil
}

func GetLogHashInteractive() (string, error) {
	logs, err := getLogs()

	if err != nil {
		return "", err
	}

	logs, err = arrangeHashPosition(logs)

	if err != nil {
		return "", err
	}

	log, err := PipeToPeco(logs)

	if err != nil {
		return "", err
	}

	return extractHash(log)
}

func getLogs() ([]string, error) {
	out, err := exec.Command("git", "log", "--oneline", "--graph", "--decorate").Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}

func execCommand(commandName string, args ...string) error {
	cmd := exec.Command(commandName, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func trimBranchName(branchName string) string {
	branchName = strings.Trim(branchName, "* \n")
	return strings.Replace(branchName, "remotes/", "", -1)
}
