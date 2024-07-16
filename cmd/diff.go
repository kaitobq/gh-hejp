/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diffOptionDescriptions = map[string]string{
	"staged": "--staged オプションは、インデックスに追加されたファイルの差分を表示します。\n使用例: git diff --staged",
	"cached": "--cached オプションは、インデックスに追加されたファイルの差分を表示します。\n使用例: git diff --cached",
	"stat":   "--stat オプションは、各ファイルの差分の統計情報を表示します。\n使用例: git diff --stat",
}

var diffLong = `diffコマンドのヘルプを表示するコマンドです。

git diffコマンドは、ワーキングディレクトリとインデックス、または2つのコミット間の差分を表示します。

基本的な使い方:
  git diff [オプション] [パス]

オプション:
  --staged    インデックスに追加されたファイルの差分を表示
  --cached    インデックスに追加されたファイルの差分を表示
  --stat      各ファイルの差分の統計情報を表示

例:
  git diff
  git diff --staged
  git diff --cached
  git diff --stat`

var diffRun = `git diff:

git diffコマンドは、ワーキングディレクトリとインデックス、または2つのコミット間の差分を表示します。

基本的な使い方:
  git diff [オプション] [パス]

オプション:
  --staged    インデックスに追加されたファイルの差分を表示
  --cached    インデックスに追加されたファイルの差分を表示
  --stat      各ファイルの差分の統計情報を表示

例:
  git diff
  git diff --staged
  git diff --cached
  git diff --stat`

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "diffコマンドのヘルプを表示するコマンドです。",
	Long:  diffLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(diffRun)
			return
		}

		for _, arg := range args {
			if desc, ok := diffOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
