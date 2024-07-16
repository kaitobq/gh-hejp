/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fetchOptionDescriptions = map[string]string{
	"all":       "--all オプションは、すべてのリモートリポジトリから変更をフェッチします。\n使用例: git fetch --all",
	"prune":     "--prune オプションは、フェッチ時にローカルで追跡されなくなったリモート追跡ブランチを削除します。\n使用例: git fetch --prune",
	"dry-run":   "--dry-run オプションは、実際にフェッチを実行せずに、どのような変更がフェッチされるかを表示します。\n使用例: git fetch --dry-run",
}

var fetchLong = `fetchコマンドのヘルプを表示するコマンドです。

git fetchコマンドは、リモートリポジトリの変更をローカルリポジトリに取り込みます。ただし、ワーキングディレクトリは更新されません。

基本的な使い方:
  git fetch <リモート> <リファレンス>

オプション:
  --all         すべてのリモートリポジトリから変更をフェッチ
  --prune       フェッチ時に不要なリモート追跡ブランチを削除
  --dry-run     実際にフェッチを実行せずに変更内容を表示

例:
  git fetch origin
  git fetch --all
  git fetch --prune
  git fetch --dry-run`

var fetchRun = `git fetch:

git fetchコマンドは、リモートリポジトリの変更をローカルリポジトリに取り込みます。ただし、ワーキングディレクトリは更新されません。

基本的な使い方:
  git fetch <リモート> <リファレンス>

オプション:
  --all         すべてのリモートリポジトリから変更をフェッチ
  --prune       フェッチ時に不要なリモート追跡ブランチを削除
  --dry-run     実際にフェッチを実行せずに変更内容を表示

例:
  git fetch origin
  git fetch --all
  git fetch --prune
  git fetch --dry-run`

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetchコマンドのヘルプを表示するコマンドです。",
	Long:  fetchLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(fetchRun)
			return
		}

		for _, arg := range args {
			if desc, ok := fetchOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
