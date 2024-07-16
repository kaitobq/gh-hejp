/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tagOptionDescriptions = map[string]string{
	"a":       "-a オプションは、注釈付きタグを作成します。注釈付きタグにはメタデータやメッセージを追加することができます。\n使用例: git tag -a v1.0.0 -m 'バージョン1.0.0'",
	"d":       "-d オプションは、指定したタグを削除します。\n使用例: git tag -d v1.0.0",
	"force":   "--force オプションは、既存のタグを強制的に上書きします。\n使用例: git tag --force v1.0.0",
	"list":    "--list オプションは、リポジトリ内のすべてのタグを一覧表示します。\n使用例: git tag --list",
}

var tagLong = `tagコマンドのヘルプを表示するコマンドです。

git tagコマンドは、特定のコミットにラベルを付けるために使用されます。タグは、特定のポイントを参照するために便利です。

基本的な使い方:
  git tag <タグ名>

オプション:
  -a, --annotate         注釈付きタグを作成
  -d, --delete           タグを削除
  --force                強制的にタグを作成または上書き
  --list                 タグを一覧表示

例:
  git tag v1.0.0
  git tag -a v1.0.0 -m 'バージョン1.0.0'
  git tag -d v1.0.0
  git tag --force v1.0.0
  git tag --list`

var tagRun = `git tag:

git tagコマンドは、特定のコミットにラベルを付けるために使用されます。タグは、特定のポイントを参照するために便利です。

基本的な使い方:
  git tag <タグ名>

オプション:
  -a, --annotate         注釈付きタグを作成
  -d, --delete           タグを削除
  --force                強制的にタグを作成または上書き
  --list                 タグを一覧表示

例:
  git tag v1.0.0
  git tag -a v1.0.0 -m 'バージョン1.0.0'
  git tag -d v1.0.0
  git tag --force v1.0.0
  git tag --list`

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "tagコマンドのヘルプを表示するコマンドです。",
	Long:  tagLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(tagRun)
			return
		}

		for _, arg := range args {
			if desc, ok := tagOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
