/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var revertOptionDescriptions = map[string]string{
	"no-commit": "--no-commit オプションは、リセットを行うが、コミットは行わずに変更をステージングエリアに残します。\n使用例: git revert --no-commit <コミットID>",
	"edit":      "--edit オプションは、コミットメッセージを編集するためのエディタを開きます。\n使用例: git revert --edit <コミットID>",
	"quit":      "--quit オプションは、現在のリバート操作を中止します。\n使用例: git revert --quit",
	"continue":  "--continue オプションは、中断されたリバート操作を再開します。\n使用例: git revert --continue",
	"abort":     "--abort オプションは、リバート操作を完全に中止し、元の状態に戻します。\n使用例: git revert --abort",
}

var revertLong = `revertコマンドのヘルプを表示するコマンドです。

git revertコマンドは、指定したコミットの変更を元に戻し、その変更を新しいコミットとして追加します。

基本的な使い方:
  git revert <コミットID>

オプション:
  --no-commit  リセットを行うが、コミットは行わない
  --edit       コミットメッセージを編集
  --quit       現在のリバート操作を中止
  --continue   中断されたリバート操作を再開
  --abort      リバート操作を完全に中止

例:
  git revert <コミットID>
  git revert --no-commit <コミットID>
  git revert --edit <コミットID>
  git revert --quit
  git revert --continue
  git revert --abort`

var revertRun = `git revert:

git revertコマンドは、指定したコミットの変更を元に戻し、その変更を新しいコミットとして追加します。

基本的な使い方:
  git revert <コミットID>

オプション:
  --no-commit  リセットを行うが、コミットは行わない
  --edit       コミットメッセージを編集
  --quit       現在のリバート操作を中止
  --continue   中断されたリバート操作を再開
  --abort      リバート操作を完全に中止

例:
  git revert <コミットID>
  git revert --no-commit <コミットID>
  git revert --edit <コミットID>
  git revert --quit
  git revert --continue
  git revert --abort`

// revertCmd represents the revert command
var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "revertコマンドのヘルプを表示するコマンドです。",
	Long:  revertLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(revertRun)
			return
		}

		for _, arg := range args {
			if desc, ok := revertOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(revertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// revertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// revertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
