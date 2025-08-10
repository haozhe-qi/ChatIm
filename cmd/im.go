package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ConfigPath string
)

func init() {
	// 1. 注册初始化函数（在 Cobra 命令执行前调用）
	cobra.OnInitialize(initConfig)

	// 2. 定义全局命令行参数 `--config`
	rootCmd.PersistentFlags().StringVar(
		&ConfigPath,                          // 绑定到变量 ConfigPath
		"config",                             // 参数名
		"./im.yaml",                          // 默认值
		"config file (default is ./im.yaml)") // 参数说明
}
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
