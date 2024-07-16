/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var optionDescriptions = map[string]string{
	"m":      "コミットメッセージを指定します。\n使用例: git commit -m \"初回コミット\"",
	"a":      "すべての変更されたファイルを自動的にステージングしてコミットします。\n使用例: git commit -a -m \"変更を追加\"",
	"amend": "最後のコミットを修正します。\n使用例: git commit --amend -m \"メッセージの修正\"",
	"dry-run": "実際にはコミットを作成せずに、何がコミットされるかを確認します。\n使用例: git commit --dry-run",
}

var long =  `commitコマンドのヘルプを表示するコマンドです。

git commitコマンドは、現在の作業ディレクトリの変更を新しいコミットとして保存します。
コミットは、プロジェクトの履歴において重要な単位であり、後で変更を元に戻すことができるポイントを作成します。

基本的な使い方:
  git commit -m "コミットメッセージ"

オプション:
  -a, --all                 すべての変更されたファイルを自動的にステージングしてコミットします
  -m, --message <メッセージ>  コミットメッセージを指定します
  --amend                  最後のコミットを修正します
  --dry-run                実際にはコミットを作成せずに、何がコミットされるかを確認します

例:
  git commit -m "初回コミット"
  git commit -a -m "変更を追加"
  git commit --amend -m "メッセージの修正"`

var run = `git commit:
		
git commitコマンドは、現在の作業ディレクトリの変更を新しいコミットとして保存します。
コミットは、プロジェクトの履歴において重要な単位であり、後で変更を元に戻すことができるポイントを作成します。

基本的な使い方:
  git commit -m "コミットメッセージ"

オプション:
  -a, --all                 すべての変更されたファイルを自動的にステージングしてコミットします
  -m, --message <メッセージ>  コミットメッセージを指定します
  --amend                  最後のコミットを修正します
  --dry-run                実際にはコミットを作成せずに、何がコミットされるかを確認します

例:
  git commit -m "初回コミット"
  git commit -a -m "変更を追加"
  git commit --amend -m "メッセージの修正"`

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commitコマンドのヘルプを表示するコマンドです。",
	Long: long,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(run)
			return
		}

		for _, arg := range args {
			if desc, ok := optionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
