//go:build windows
// +build windows

package go_shell

import (
	"os/exec"
)

func NewCmd(s string) *exec.Cmd {
	cmd := exec.Command("cmd", "/C", s)
	return cmd
}
