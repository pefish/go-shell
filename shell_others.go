//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package go_shell

var ErrNotImplemented = errors.New("shell: not implemented")

func NewCmd(s string) *exec.Cmd {
	panic(ErrNotImplemented)
}
