package gico

import (
	"strings"
	"os/exec"
	"os"
	"io"
)

func SwitchBranch() error {
	names, err := getBranchNames()

	str, err := pipeToPeco(names)

	if err != nil {
		return err
	}

	if len(str) == 0 {
		return nil
	}

	branchName := strings.Trim(string(str), " \n")
	err = execCommand("git", "checkout", branchName)

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

func pipeToPeco(texts []string) (string, error) {

	cmd := exec.Command("peco")
	stdin, _ := cmd.StdinPipe()
	io.WriteString(stdin, strings.Join(texts, "\n"))
	stdin.Close()
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
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