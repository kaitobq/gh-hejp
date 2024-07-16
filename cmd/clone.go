/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cloneOptionDescriptions = map[string]string{
	"bare": "--bare オプションは、ワーキングディレクトリを作成せずにリポジトリをクローンします。\n使用例: git clone --bare <リポジトリURL>",
	"depth": "--depth オプションは、指定した数のコミット履歴だけを取得します。\n使用例: git clone --depth 1 <リポジトリURL>",
	"branch": "--branch オプションは、指定したブランチだけをクローンします。\n使用例: git clone --branch <ブランチ名> <リポジトリURL>",
}

var cloneLong = `cloneコマンドのヘルプを表示するコマンドです。

git cloneコマンドは、リモートリポジトリのコピーをローカルに作成します。

基本的な使い方:
  git clone <リポジトリURL>

オプション:
  --bare                 ワーキングディレクトリを作成せずにクローン
  --depth <数>           指定した数のコミット履歴だけを取得
  --branch <ブランチ名>  指定したブランチだけをクローン

例:
  git clone https://github.com/user/repo.git
  git clone --bare https://github.com/user/repo.git
  git clone --depth 1 https://github.com/user/repo.git
  git clone --branch develop https://github.com/user/repo.git`

var cloneRun = `git clone:

git cloneコマンドは、リモートリポジトリのコピーをローカルに作成します。

基本的な使い方:
  git clone <リポジトリURL>

オプション:
  --bare                 ワーキングディレクトリを作成せずにクローン
  --depth <数>           指定した数のコミット履歴だけを取得
  --branch <ブランチ名>  指定したブランチだけをクローン

例:
  git clone https://github.com/user/repo.git
  git clone --bare https://github.com/user/repo.git
  git clone --depth 1 https://github.com/user/repo.git
  git clone --branch develop https://github.com/user/repo.git`

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "cloneコマンドのヘルプを表示するコマンドです。",
	Long:  cloneLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cloneRun)
			return
		}

		for _, arg := range args {
			if desc, ok := cloneOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
