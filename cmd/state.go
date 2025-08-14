package cmd

import (
	"github.com/haozhe-qi/ChatIm/state"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stateCmd)
}

var stateCmd = &cobra.Command{
	Use: "state",
	Run: StateHandler,
}

func StateHandler(cmd *cobra.Command, args []string) {
	state.RunMain(ConfigPath)
}
