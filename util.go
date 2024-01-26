package go_shell

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
	"time"
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

func ExecForResult(cmd *exec.Cmd) (string, error) {
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func ExecForResultLineByLine(cmd *exec.Cmd, resultChan chan<- string) error {
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

	errMsgs := make([]string, 0)
	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			errMsgs = append(errMsgs, errScanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		return errors.New(strings.Join(errMsgs, "\n"))
	}

	return nil
}
