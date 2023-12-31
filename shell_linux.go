//go:build linux
// +build linux

package go_shell

import (
	"os"
	"os/exec"
)

func NewCmd(s string) *exec.Cmd {
	cmd := exec.Command("/bin/bash", "-c", s)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
