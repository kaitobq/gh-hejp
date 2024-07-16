/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmOptionDescriptions = map[string]string{
	"f": "-f オプションは、警告なしで強制的にファイルを削除します。\n使用例: git rm -f <ファイル>",
	"r": "-r オプションは、再帰的にディレクトリを削除します。\n使用例: git rm -r <ディレクトリ>",
	"cached": "--cached オプションは、インデックスからファイルを削除しますが、作業ディレクトリからは削除しません。\n使用例: git rm --cached <ファイル>",
}

var rmLong = `rmコマンドのヘルプを表示するコマンドです。

git rmコマンドは、インデックスと作業ディレクトリからファイルを削除します。

基本的な使い方:
  git rm <ファイル>

オプション:
  -f, --force       強制的にファイルを削除
  -r, --recursive   再帰的にディレクトリを削除
  --cached          インデックスから削除

例:
  git rm file.txt
  git rm -f file.txt
  git rm -r directory
  git rm --cached file.txt`

var rmRun = `git rm:

git rmコマンドは、インデックスと作業ディレクトリからファイルを削除します。

基本的な使い方:
  git rm <ファイル>

オプション:
  -f, --force       強制的にファイルを削除
  -r, --recursive   再帰的にディレクトリを削除
  --cached          インデックスから削除

例:
  git rm file.txt
  git rm -f file.txt
  git rm -r directory
  git rm --cached file.txt`

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "rmコマンドのヘルプを表示するコマンドです。",
	Long:  rmLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(rmRun)
			return
		}

		for _, arg := range args {
			if desc, ok := rmOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
