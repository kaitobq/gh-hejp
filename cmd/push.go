/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushOptionDescriptions = map[string]string{
	"u":       "-u オプションは、現在のブランチを指定したリモートブランチにアップストリームとして設定します。\n使用例: git push -u origin ブランチ名",
	"force":  "--force オプションは、リモートリポジトリの履歴を上書きする強制プッシュを行います。\n使用例: git push --force origin ブランチ名",
	"tags":   "--tags オプションは、タグを含めてすべてのリファレンスをプッシュします。\n使用例: git push --tags",
}

var pushLong = `pushコマンドのヘルプを表示するコマンドです。

git pushコマンドは、ローカルリポジトリの変更をリモートリポジトリに送信します。

基本的な使い方:
  git push <リモート> <ブランチ>

オプション:
  -u, --set-upstream    アップストリームとして設定
  --force               強制的にプッシュ
  --tags                タグをプッシュ

例:
  git push origin main
  git push -u origin feature/new-feature
  git push --force origin main
  git push --tags`

var pushRun = `git push:

git pushコマンドは、ローカルリポジトリの変更をリモートリポジトリに送信します。

基本的な使い方:
  git push <リモート> <ブランチ>

オプション:
  -u, --set-upstream    アップストリームとして設定
  --force               強制的にプッシュ
  --tags                タグをプッシュ

例:
  git push origin main
  git push -u origin feature/new-feature
  git push --force origin main
  git push --tags`

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "pushコマンドのヘルプを表示するコマンドです。",
	Long:  pushLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(pushRun)
			return
		}

		for _, arg := range args {
			if desc, ok := pushOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
