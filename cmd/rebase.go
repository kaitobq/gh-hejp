/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rebaseOptionDescriptions = map[string]string{
	"i":       "-i オプションは、対話的にリベースを行います。\n使用例: git rebase -i <ブランチ>",
	"abort":   "--abort オプションは、進行中のリベースを中止し、リベース前の状態に戻します。\n使用例: git rebase --abort",
	"continue": "--continue オプションは、停止したリベースを続行します。\n使用例: git rebase --continue",
}

var rebaseLong = `rebaseコマンドのヘルプを表示するコマンドです。

git rebaseコマンドは、あるブランチの変更を別のブランチの上に適用します。リベースにより、履歴を整理し、リニアな履歴を保つことができます。

基本的な使い方:
  git rebase <ブランチ>

オプション:
  -i, --interactive    対話的にリベースを行います
  --abort              進行中のリベースを中止します
  --continue           停止したリベースを続行します

例:
  git rebase main
  git rebase -i main
  git rebase --abort
  git rebase --continue`

var rebaseRun = `git rebase:

git rebaseコマンドは、あるブランチの変更を別のブランチの上に適用します。リベースにより、履歴を整理し、リニアな履歴を保つことができます。

基本的な使い方:
  git rebase <ブランチ>

オプション:
  -i, --interactive    対話的にリベースを行います
  --abort              進行中のリベースを中止します
  --continue           停止したリベースを続行します

例:
  git rebase main
  git rebase -i main
  git rebase --abort
  git rebase --continue`

// rebaseCmd represents the rebase command
var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "rebaseコマンドのヘルプを表示するコマンドです。",
	Long:  rebaseLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(rebaseRun)
			return
		}

		for _, arg := range args {
			if desc, ok := rebaseOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}


func init() {
	rootCmd.AddCommand(rebaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rebaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rebaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
