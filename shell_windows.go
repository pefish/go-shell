//go:build windows
// +build windows

package go_shell

import (
	"fmt"
	"os/exec"
)

func NewCmd(format string, a ...any) *exec.Cmd {
	cmd := exec.Command("cmd", "/C", fmt.Sprintf(format, a...))
	return cmd
}
