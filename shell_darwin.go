//go:build darwin
// +build darwin

package go_shell

import (
	"os"
	"os/exec"
)

func NewCmd(s string) *exec.Cmd {
	cmd := exec.Command("/cmd/bash", "-c", s)
	cmd.Stdout = os.Stdout
	return cmd
}
