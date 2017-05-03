package utils

import (
	"os/exec"
	"io"
	"strings"
	"errors"
	"regexp"
	"os"
	"github.com/mpppk/gico/etc"
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

func ParseRemoteURL(url string, hosts []*etc.Host) (host, owner, repoName string, err error) {

	var hostName string
	for _, host := range hosts {
		if strings.Contains(url, host.Host) {
			hostName = host.Host
		}
	}

	assigned := regexp.MustCompile( strings.Replace(hostName, ".", `\.`, -1) + `.(.*)`)
	group := assigned.FindStringSubmatch(url)

	if len(group) < 2 {
		return "", "", "", errors.New("invalid url: " + url)
	}

	ownerAndRepo := strings.Split(group[1], "/")
	repoName = strings.Replace(ownerAndRepo[1], ".git", "", -1)
	return strings.Replace(hostName, ".com", "", -1), ownerAndRepo[0], repoName, nil
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
