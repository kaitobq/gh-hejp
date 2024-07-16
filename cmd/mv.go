/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mvOptionDescriptions = map[string]string{
	"-f": "-f オプションは、上書き確認なしでファイルを強制的に移動します。\n使用例: git mv -f <旧パス> <新パス>",
}

var mvLong = `mvコマンドのヘルプを表示するコマンドです。

git mvコマンドは、ファイルまたはディレクトリを移動またはリネームします。

基本的な使い方:
  git mv <旧パス> <新パス>

オプション:
  -f  強制的に移動またはリネーム

例:
  git mv README.md README.old.md
  git mv -f file1.txt file2.txt`

var mvRun = `git mv:

git mvコマンドは、ファイルまたはディレクトリを移動またはリネームします。

基本的な使い方:
  git mv <旧パス> <新パス>

オプション:
  -f  強制的に移動またはリネーム

例:
  git mv README.md README.old.md
  git mv -f file1.txt file2.txt`

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "mvコマンドのヘルプを表示するコマンドです。",
	Long:  mvLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(mvRun)
			return
		}

		for _, arg := range args {
			if desc, ok := mvOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
