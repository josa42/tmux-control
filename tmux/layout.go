package tmux

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadLayout(name string) Layout {
	layout := Layout{}

	home, _ := os.UserHomeDir()
	lpath := filepath.Join(home, ".config", "tmux-control", "layouts", fmt.Sprintf("%s.yml", name))

	data, _ := ioutil.ReadFile(lpath)

	err := yaml.Unmarshal(data, &layout)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	return layout
}

// Example:
//
// ```yaml
// windows:
//   - name: main
//     layout: main-vertical
//     panes:
//       - run: nvim
//       - run: ls
//       - run: ls
//   - name: second
//     layout: main-horizontal
//     panes:
//       - run: tail -f /var/log/system.log
//       - run: ls
//       - run: ls
//   - name: third
//     panes:
//       - run: tree
// ```

type Layout struct {
	Windows []Window
}

type Window struct {
	Name   string
	Layout string
	Panes  []Pane
}

type Pane struct {
	Name string
	Run  string
}
