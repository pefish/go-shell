//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package go_shell

import (
	"errors"
	"os/exec"
)

var ErrNotImplemented = errors.New("shell: not implemented")

func NewCmd(format string, a ...any) *exec.Cmd {
	panic(ErrNotImplemented)
}
