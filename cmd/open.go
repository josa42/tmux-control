package cmd

import (
	"github.com/josa42/tmux-control/tmux"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <name>",
	Short: "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if !tmux.SessionExists(name) {
			tmux.NewSession(name)
		}
		tmux.FocusSession(name)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
