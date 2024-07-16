/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configOptionDescriptions = map[string]string{
	"list":  "--list オプションは、すべての設定を表示します。\n使用例: git config --list",
	"global": "--global オプションは、グローバル設定に変更を加えます。\n使用例: git config --global user.name 'Your Name'",
	"system": "--system オプションは、システム全体の設定に変更を加えます。\n使用例: git config --system core.editor vim",
}

var configLong = `configコマンドのヘルプを表示するコマンドです。

git configコマンドは、Gitの設定を管理します。

基本的な使い方:
  git config [オプション] <キー> <値>

オプション:
  --list                 すべての設定を表示
  --global               グローバル設定に変更を加える
  --system               システム全体の設定に変更を加える

例:
  git config --list
  git config --global user.name 'Your Name'
  git config --system core.editor vim`

var configRun = `git config:

git configコマンドは、Gitの設定を管理します。

基本的な使い方:
  git config [オプション] <キー> <値>

オプション:
  --list                 すべての設定を表示
  --global               グローバル設定に変更を加える
  --system               システム全体の設定に変更を加える

例:
  git config --list
  git config --global user.name 'Your Name'
  git config --system core.editor vim`

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configコマンドのヘルプを表示するコマンドです。",
	Long:  configLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(configRun)
			return
		}

		for _, arg := range args {
			if desc, ok := configOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
