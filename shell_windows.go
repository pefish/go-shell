//go:build windows
// +build windows

package go_shell

import (
	"os"
	"os/exec"
)

func NewCmd(s string) *exec.Cmd {
	cmd := exec.Command("cmd", "/C", s)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
