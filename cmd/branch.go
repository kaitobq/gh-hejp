/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var branchOptionDescriptions = map[string]string{
	"a":   "--all オプションは、リモート追跡ブランチを含むすべてのブランチを表示します。\n使用例: git branch -a",
	"d":   "--delete オプションは、指定したブランチを削除します。\n使用例: git branch -d branch_name",
	"m":   "--move オプションは、ブランチの名前を変更します。\n使用例: git branch -m old_name new_name",
	"list": "--list オプションは、指定したパターンに一致するブランチをリストします。\n使用例: git branch --list 'feature/*'",
}

var branchLong = `branchコマンドのヘルプを表示するコマンドです。

git branchコマンドは、新しいブランチの作成、既存のブランチの削除、ブランチの名前変更、現在のブランチの表示など、ブランチ操作を行います。

基本的な使い方:
  git branch

オプション:
  -a, --all           すべてのブランチを表示します
  -d, --delete        指定したブランチを削除します
  -m, --move          ブランチの名前を変更します
  --list              指定したパターンに一致するブランチをリストします

例:
  git branch
  git branch -a
  git branch -d branch_name
  git branch -m old_name new_name
  git branch --list 'feature/*'`

var branchRun = `git branch:

git branchコマンドは、新しいブランチの作成、既存のブランチの削除、ブランチの名前変更、現在のブランチの表示など、ブランチ操作を行います。

基本的な使い方:
  git branch

オプション:
  -a, --all           すべてのブランチを表示します
  -d, --delete        指定したブランチを削除します
  -m, --move          ブランチの名前を変更します
  --list              指定したパターンに一致するブランチをリストします

例:
  git branch
  git branch -a
  git branch -d branch_name
  git branch -m old_name new_name
  git branch --list 'feature/*'`

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "branchコマンドのヘルプを表示するコマンドです。",
	Long:  branchLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(branchRun)
			return
		}

		for _, arg := range args {
			if desc, ok := branchOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
