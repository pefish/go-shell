package go_shell

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
)

func ExecInConsole(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func ExecResult(cmd *exec.Cmd) (string, error) {
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func ExecResultLineByLine(cmd *exec.Cmd, resultChan chan<- string) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			resultChan <- scanner.Text()
		}
	}()

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
