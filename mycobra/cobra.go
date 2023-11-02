package mycobra

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var a bool
var b bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your application",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api called")
		a, _ = cmd.Flags().GetBool("toggle")
		b, _ = cmd.Flags().GetBool("out")
		fmt.Println(a, b)
	},
}

// 命令一
var mockMsgCmd = &cobra.Command{
	Use:   "mockMsg",
	Short: "批量发送测试文本消息",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mockMsg called")
		fmt.Println(args)
	},
}

// 命令二
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "导出数据",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
		out, _ := cmd.Flags().GetString("out")
		cfgfile, _ := cmd.Flags().GetString("cfgfile")
		fmt.Println(args)
		fmt.Println(out, cfgfile)
	},
}

var testVar bool
var cfgFile string

func init() {

	// Flags 是 当前命令行（rootCmd 或 exportCmd）的参数。
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVar(&testVar, "paramName", true, "这里是我测试的参数")
	exportCmd.Flags().StringP("out", "k", "./backup", "导出路径")

	// PersistentFlags 是 当前命令行和子命令行内都可以获取到这个参数。
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "cfgfile", "c", ".env", "配置文件路径和文件")

	// mockMsgCmd 和 exportCmd 被添加到rootCmd的子命令，
	rootCmd.AddCommand(mockMsgCmd)
	rootCmd.AddCommand(exportCmd)

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(a, b)
}
