package utils

import (
	"os/exec"
	"io"
	"strings"
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
