/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addOptionDescriptions = map[string]string{
	"A":    "ワーキングツリー内の全ての変更をステージングします。\n使用例: git add -A",
	"n":    "実際にはステージングせず、何がステージングされるかを確認します。\n使用例: git add -n .",
	"v":    "ステージングするファイルの詳細を表示します。\n使用例: git add -v .",
	"all": "ワーキングツリー内の全ての変更をステージングします（-Aと同じ）。\n使用例: git add --all",
}

var addLong = `addコマンドのヘルプを表示するコマンドです。

git addコマンドは、指定したファイルまたはディレクトリをステージングエリアに追加します。
ステージングエリアに追加された変更は、次のコミットに含まれます。

基本的な使い方:
  git add <ファイル名/ディレクトリ名>

オプション:
  -A, --all        ワーキングツリー内の全ての変更をステージングします
  -n, --dry-run    実際にはステージングせず、何がステージングされるかを確認します
  -v, --verbose    ステージングするファイルの詳細を表示します

例:
  git add .
  git add -A
  git add -n .
  git add -v .`

var addRun = `git add:

git addコマンドは、指定したファイルまたはディレクトリをステージングエリアに追加します。
ステージングエリアに追加された変更は、次のコミットに含まれます。

基本的な使い方:
  git add <ファイル名/ディレクトリ名>

オプション:
  -A, --all        ワーキングツリー内の全ての変更をステージングします
  -n, --dry-run    実際にはステージングせず、何がステージングされるかを確認します
  -v, --verbose    ステージングするファイルの詳細を表示します

例:
  git add .
  git add -A
  git add -n .
  git add -v .`

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "addコマンドのヘルプを表示するコマンドです。",
	Long:  addLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(addRun)
			return
		}

		for _, arg := range args {
			if desc, ok := addOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
