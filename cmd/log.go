/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logOptionDescriptions = map[string]string{
	"oneline": "--oneline オプションは、各コミットを1行で表示します。\n使用例: git log --oneline",
	"graph":   "--graph オプションは、ブランチとマージの履歴をASCIIグラフで表示します。\n使用例: git log --graph",
	"decorate": "--decorate オプションは、リファレンス名を表示します。\n使用例: git log --decorate",
}

var logLong = `logコマンドのヘルプを表示するコマンドです。

git logコマンドは、リポジトリのコミット履歴を表示します。

基本的な使い方:
  git log

オプション:
  --oneline             各コミットを1行で表示
  --graph               ブランチとマージの履歴をASCIIグラフで表示
  --decorate            リファレンス名を表示

例:
  git log
  git log --oneline
  git log --graph
  git log --decorate`

var logRun = `git log:

git logコマンドは、リポジトリのコミット履歴を表示します。

基本的な使い方:
  git log

オプション:
  --oneline             各コミットを1行で表示
  --graph               ブランチとマージの履歴をASCIIグラフで表示
  --decorate            リファレンス名を表示

例:
  git log
  git log --oneline
  git log --graph
  git log --decorate`

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "logコマンドのヘルプを表示するコマンドです。",
	Long:  logLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(logRun)
			return
		}

		for _, arg := range args {
			if desc, ok := logOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
