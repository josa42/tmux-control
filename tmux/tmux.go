package tmux

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ApplySessionLayout(name string, layout string) {
	if !SessionIsEmpty(name) {
		fmt.Println("error: layout can only be applied to empty sessions")
		os.Exit(1)
	}

	l := LoadLayout(layout)
	windows := execTmuxList("list-windows", "-t", name, "-F", "#{window_id}")
	for idx, w := range l.Windows {
		wname := w.Name
		if wname == "" {
			wname = fmt.Sprintf("%d", idx)
		}

		if idx == 0 {
			RenameWindow(windows[0], wname)
		} else {
			NewWindow(name, wname)
		}
	}

	windows = execTmuxList("list-windows", "-t", name, "-F", "#{window_id}")

	for widx, w := range l.Windows {

		for pidx := range w.Panes {
			if pidx != 0 {
				panes := execTmuxList("list-panes", "-t", windows[widx], "-F", "#{pane_id}")
				execTmux("split-window", "-t", panes[pidx-1])
			}
		}

		panes := execTmuxList("list-panes", "-t", windows[widx], "-F", "#{pane_id}")

		if w.Layout != "" {
			execTmux("select-layout", "-t", panes[0], w.Layout)
		}

		for pidx, p := range w.Panes {
			if p.Run != "" {
				SendKeys(panes[pidx], p.Run)
			}
		}
	}
}

func SessionExists(name string) bool {
	out := execTmux("list-sessions", "-F", "#{session_name}")
	for _, n := range strings.Split(out, "\n") {
		if n == name {
			return true
		}
	}

	return false
}

func SessionIsEmpty(name string) bool {
	windows := execTmuxList("list-windows", "-t", name, "-F", "#{window_id}")
	panes := execTmuxList("list-panes", "-t", name, "-F", "#{pane_id}")

	return len(windows) == 1 && len(panes) == 1
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

func ChangeDirectory(name string, directory string) {
	SendKeys(name, fmt.Sprintf("cd '%s'", directory))
}

func Clear(name string) {
	SendKeys(name, "clear; tmux clear-history; clear")
}

func SendKeys(name string, command string) {
	execTmux("send-keys", "-t", name, command, "C-m")
}

func NewWindow(sessionName string, name string) {
	execTmux("new-window", "-t", sessionName, "-n", name, "-d")
}

func RenameWindow(window, name string) {
	execTmux("rename-window", "-t", window, name)
}

func execTmux(args ...string) string {
	cmd := exec.Command("tmux", args...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()

	if err != nil {
		fmt.Printf("error: %v", errb.String())
		os.Exit(1)
	}

	return outb.String()
}

func execTmuxList(args ...string) []string {
	out := execTmux(args...)

	return strings.Split(strings.Trim(out, "\n"), "\n")
}

type ExecError struct {
	err error
	Out string
}

func (e *ExecError) Error() string {
	return e.err.Error()
}
