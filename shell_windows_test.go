//go:build windows
// +build windows

package go_shell

import (
	"testing"

	go_test_ "github.com/pefish/go-test"
)

func TestNewCmd(t *testing.T) {
	result, err := ExecResult(NewCmd("ls"))
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, true, len(result) > 0)
}
