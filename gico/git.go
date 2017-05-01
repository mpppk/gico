package gico

import (
	"os/exec"
	"regexp"
	"strings"
)

func GetBranchNames() ([]string, error) {
	out, err := exec.Command("git", "branch", "-a").Output()

	if err != nil {
		return nil, err
	}

	branchNames := strings.Split(string(out), "\n")

	return branchNames, nil
}

func ExtractHash(log string) (string, error) {
	reg := regexp.MustCompile(`[0123456789abcdef]{7}`)
	return reg.FindString(log), nil
}

func ArrangeHashPosition(logs []string) ([]string, error) {
	var newLogs []string
	for _, log := range logs {
		hash, err := ExtractHash(log)

		if err != nil {
			return nil, err
		}

		newLog := strings.Replace(log, hash, "", -1)

		newLogs = append(newLogs, newLog+" ["+hash+"]")
	}
	return newLogs, nil
}

func GetLogs() ([]string, error) {
	out, err := exec.Command("git", "log", "--oneline", "--graph", "--decorate").Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}


func TrimBranchName(branchName string) string {
	branchName = strings.Trim(branchName, "* \n")
	return strings.Replace(branchName, "remotes/", "", -1)
}
