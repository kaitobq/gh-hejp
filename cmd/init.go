/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initOptionDescriptions = map[string]string{
	"bare": "--bare オプションは、ベアリポジトリを作成します。ベアリポジトリは作業ツリーを持たず、リモートリポジトリとして使用されます。\n使用例: git init --bare",
}

var initLong = `initコマンドのヘルプを表示するコマンドです。

git initコマンドは、新しいGitリポジトリを作成します。
このコマンドを実行すると、.gitディレクトリが作成され、そこにGitリポジトリに関するすべてのメタデータが保存されます。

基本的な使い方:
  git init [オプション] [ディレクトリ]

オプション:
  --bare  ベアリポジトリとして初期化します

例:
  git init
  git init --bare
  git init my-project`

var initRun = `git init:

git initコマンドは、新しいGitリポジトリを作成します。
このコマンドを実行すると、.gitディレクトリが作成され、そこにGitリポジトリに関するすべてのメタデータが保存されます。

基本的な使い方:
  git init [オプション] [ディレクトリ]

オプション:
  --bare  ベアリポジトリとして初期化します

例:
  git init
  git init --bare
  git init my-project`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initコマンドのヘルプを表示するコマンドです。",
	Long:  initLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(initRun)
			return
		}

		for _, arg := range args {
			if desc, ok := initOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// initCmd.Flags().BoolVar(&bare, "bare", false, "ベアリポジトリとして初期化します")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
