package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use:   "ChatIm",
	Short: "这是一个IM系统",
	Run:   Im,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func Im(cmd *cobra.Command, args []string) {

}
func initConfig() {

}
