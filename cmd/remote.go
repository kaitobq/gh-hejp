/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var remoteOptionDescriptions = map[string]string{
	"add":       "add オプションは、新しいリモートリポジトリを追加します。\n使用例: git remote add <リモート名> <リポジトリURL>",
	"remove":    "remove オプションは、既存のリモートリポジトリを削除します。\n使用例: git remote remove <リモート名>",
	"show":      "show オプションは、リモートリポジトリの詳細を表示します。\n使用例: git remote show <リモート名>",
	"rename":    "rename オプションは、既存のリモートリポジトリの名前を変更します。\n使用例: git remote rename <旧リモート名> <新リモート名>",
}

var remoteLong = `remoteコマンドのヘルプを表示するコマンドです。

git remoteコマンドは、リモートリポジトリの管理を行います。リモートリポジトリは、複数の開発者が共同で作業する際に使用されます。

基本的な使い方:
  git remote <オプション> <引数>

オプション:
  add       新しいリモートリポジトリを追加します
  remove    既存のリモートリポジトリを削除します
  show      リモートリポジトリの詳細を表示します
  rename    既存のリモートリポジトリの名前を変更します

例:
  git remote add origin https://github.com/user/repo.git
  git remote remove origin
  git remote show origin
  git remote rename origin upstream`

var remoteRun = `git remote:

git remoteコマンドは、リモートリポジトリの管理を行います。リモートリポジトリは、複数の開発者が共同で作業する際に使用されます。

基本的な使い方:
  git remote <オプション> <引数>

オプション:
  add       新しいリモートリポジトリを追加します
  remove    既存のリモートリポジトリを削除します
  show      リモートリポジトリの詳細を表示します
  rename    既存のリモートリポジトリの名前を変更します

例:
  git remote add origin https://github.com/user/repo.git
  git remote remove origin
  git remote show origin
  git remote rename origin upstream`

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "remoteコマンドのヘルプを表示するコマンドです。",
	Long:  remoteLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(remoteRun)
			return
		}

		for _, arg := range args {
			if desc, ok := remoteOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
