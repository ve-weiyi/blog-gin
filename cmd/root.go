package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-gin/cmd/migrate"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "blog-gin",
		Short:   "ve-blog-gin 博客系统",
		Version: fmt.Sprintf("%s %s/%s", "v1.0.0", runtime.GOOS, runtime.GOARCH),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("欢迎使用 ve-blog-gin 博客系统！")
			fmt.Printf("版本: %s\n", cmd.Version)
			fmt.Printf("使用 --help 查看可用命令和选项。\n")
		},
	}

	rootCmd.AddCommand(migrate.NewMigrateCmd())

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	if err := NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
