package cmd

import (
	"github.com/haozhe-qi/ChatIm/gateway"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
	Use: "gateway",
	Run: GateewayHandler,
}

func GateewayHandler(cmd *cobra.Command, args []string) {
	gateway.Runmain(ConfigPath)
}
