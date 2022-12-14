package tmux

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SessionExists(name string) bool {
	out := execTmux("list-sessions", "-F", "#{session_name}")
	for _, n := range strings.Split(out, "\n") {
		if n == name {
			return true
		}
	}

	return false
}

func FocusSession(name string) {
	execTmux("switch-client", "-t", name)
}

func NewSession(name string) {
	execTmux("new-session", "-s", name, "-d")
}

func KillSession(name string) {
	execTmux("kill-session", "-t", name)
}

func execTmux(arg ...string) string {
	cmd := exec.Command("tmux", arg...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()

	if err != nil {
		fmt.Printf("error: %v", outb.String())
		os.Exit(1)
	}

	return outb.String()
}

type ExecError struct {
	err error
	Out string
}

func (e *ExecError) Error() string {
	return e.err.Error()
}
