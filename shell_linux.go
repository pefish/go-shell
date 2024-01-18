//go:build linux
// +build linux

package go_shell

import (
	"os/exec"
)

func NewCmd(s string) *exec.Cmd {
	cmd := exec.Command("/bin/bash", "-c", s)
	return cmd
}
