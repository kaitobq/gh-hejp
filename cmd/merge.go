/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mergeOptionDescriptions = map[string]string{
	"no-ff": "--no-ff オプションは、マージを行うときに常に新しいコミットを作成します。\n使用例: git merge --no-ff ブランチ名",
	"squash": "--squash オプションは、すべてのコミットを1つの新しいコミットにまとめてマージします。\n使用例: git merge --squash ブランチ名",
	"abort": "--abort オプションは、マージが競合した場合にマージ操作を中止します。\n使用例: git merge --abort",
}

var mergeLong = `mergeコマンドのヘルプを表示するコマンドです。

git mergeコマンドは、他のブランチを現在のブランチに統合します。

基本的な使い方:
  git merge <ブランチ>

オプション:
  --no-ff                常に新しいコミットを作成
  --squash               すべてのコミットを1つにまとめてマージ
  --abort                マージ操作を中止

例:
  git merge feature-branch
  git merge --no-ff feature-branch
  git merge --squash feature-branch
  git merge --abort`

var mergeRun = `git merge:

git mergeコマンドは、他のブランチを現在のブランチに統合します。

基本的な使い方:
  git merge <ブランチ>

オプション:
  --no-ff                常に新しいコミットを作成
  --squash               すべてのコミットを1つにまとめてマージ
  --abort                マージ操作を中止

例:
  git merge feature-branch
  git merge --no-ff feature-branch
  git merge --squash feature-branch
  git merge --abort`

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "mergeコマンドのヘルプを表示するコマンドです。",
	Long:  mergeLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(mergeRun)
			return
		}

		for _, arg := range args {
			if desc, ok := mergeOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
