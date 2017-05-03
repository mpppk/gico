package utils

import (
	"os/exec"
	"io"
	"strings"
	"errors"
	"regexp"
	"os"
)

func PipeToPeco(texts []string) (string, error) {
	cmd := exec.Command("peco")
	stdin, _ := cmd.StdinPipe()
	io.WriteString(stdin, strings.Join(texts, "\n"))
	stdin.Close()
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.Trim(string(out), " \n"), nil
}

func PanicIfErrorExist(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseRemoteURL(url string) (host, owner, repoName string, err error) {

	var provider string

	if strings.Contains(url, "github.com") {
		provider = "github.com"
	}else if strings.Contains(url, "gitlab.com") {
		provider = "gitlab.com"
	}else {
		return "", "", "", errors.New("unknown host: " + url)
	}

	assined := regexp.MustCompile( strings.Replace(provider, ".", `\.`, -1) + `.(.*)`)
	group := assined.FindStringSubmatch(url)

	if len(group) < 2 {
		return "", "", "", errors.New("invalid url: " + url)
	}

	ownerAndRepo := strings.Split(group[1], "/")
	repoName = strings.Replace(ownerAndRepo[1], ".git", "", -1)
	return strings.Replace(provider, ".com", "", -1), ownerAndRepo[0], repoName, nil
}

func ExecCommand(commandName string, args ...string) error {
	cmd := exec.Command(commandName, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
