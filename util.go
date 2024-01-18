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
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	outScanner := bufio.NewScanner(stdout)
	go func() {
		for outScanner.Scan() {
			resultChan <- outScanner.Text()
		}
	}()

	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			resultChan <- errScanner.Text()
		}
	}()

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
