//go:build linux
// +build linux

package go_shell

import (
	"fmt"
	"os/exec"
)

func NewCmd(format string, a ...any) *exec.Cmd {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf(format, a...))
	return cmd
}
