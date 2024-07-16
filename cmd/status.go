/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusOptionDescriptions = map[string]string{
	"s": "--short オプションは、出力を短くします。\n使用例: git status -s",
	"b": "--branch オプションは、現在のブランチを表示します。\n使用例: git status -b",
}

var statusLong = `statusコマンドのヘルプを表示するコマンドです。

git statusコマンドは、作業ディレクトリの状態を表示します。
現在のブランチ、追跡しているリモートブランチとの違い、ステージングされた変更、ステージングされていない変更、
および追跡されていないファイルなどの情報を提供します。

基本的な使い方:
  git status

オプション:
  -s, --short   出力を短く表示します
  -b, --branch  現在のブランチを表示します

例:
  git status
  git status -s
  git status -b`

var statusRun = `git status:

git statusコマンドは、作業ディレクトリの状態を表示します。
現在のブランチ、追跡しているリモートブランチとの違い、ステージングされた変更、ステージングされていない変更、
および追跡されていないファイルなどの情報を提供します。

基本的な使い方:
  git status

オプション:
  -s, --short   出力を短く表示します
  -b, --branch  現在のブランチを表示します

例:
  git status
  git status -s
  git status -b`

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "statusコマンドのヘルプを表示するコマンドです。",
	Long:  statusLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(statusRun)
			return
		}

		for _, arg := range args {
			if desc, ok := statusOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
