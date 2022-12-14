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
		force, _ := cmd.Flags().GetBool("force")
		dir, _ := cmd.Flags().GetString("dir")
		run, _ := cmd.Flags().GetString("run")

		name := args[0]
		exists := tmux.SessionExists(name)

		if force && exists {
			tmux.KillSession(name)
		}

		if !exists || force {
			tmux.NewSession(name)
			if dir != "" {
				tmux.ChangeDirectory(name, dir)
				tmux.Clear(name)
			}
		}

		if run != "" {
			tmux.SendKeys(name, run)
		}

		tmux.FocusSession(name)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.Flags().StringP("run", "r", "", "Run command on creating")
	openCmd.Flags().StringP("dir", "d", "", "Change into directory on creating")
	openCmd.Flags().BoolP("force", "f", false, "Kill existing session")
}
