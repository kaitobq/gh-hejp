/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkoutOptionDescriptions = map[string]string{
	"b":      "-b オプションは、新しいブランチを作成してそのブランチに切り替えます。\n使用例: git checkout -b 新しいブランチ名",
	"track": "--track オプションは、リモートブランチを追跡する新しいローカルブランチを作成します。\n使用例: git checkout --track origin/ブランチ名",
}

var checkoutLong = `checkoutコマンドのヘルプを表示するコマンドです。

git checkoutコマンドは、ブランチを切り替えるか、特定のファイルをインデックスまたは作業ツリーにチェックアウトします。

基本的な使い方:
  git checkout <branch>
  git checkout <commit>
  git checkout <file>

オプション:
  -b      新しいブランチを作成してそのブランチに切り替えます
  --track リモートブランチを追跡する新しいローカルブランチを作成します

例:
  git checkout -b feature/new-feature
  git checkout --track origin/feature/new-feature
  git checkout main
  git checkout 5d41402abc4b2a76b9719d911017c592
  git checkout HEAD^ -- <file>`

var checkoutRun = `git checkout:

git checkoutコマンドは、ブランチを切り替えるか、特定のファイルをインデックスまたは作業ツリーにチェックアウトします。

基本的な使い方:
  git checkout <branch>
  git checkout <commit>
  git checkout <file>

オプション:
  -b      新しいブランチを作成してそのブランチに切り替えます
  --track リモートブランチを追跡する新しいローカルブランチを作成します

例:
  git checkout -b feature/new-feature
  git checkout --track origin/feature/new-feature
  git checkout main
  git checkout 5d41402abc4b2a76b9719d911017c592
  git checkout HEAD^ -- <file>`

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkoutコマンドのヘルプを表示するコマンドです。",
	Long:  checkoutLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(checkoutRun)
			return
		}

		for _, arg := range args {
			if desc, ok := checkoutOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
