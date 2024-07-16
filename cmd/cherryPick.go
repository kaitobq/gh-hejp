/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cherryPickOptionDescriptions = map[string]string{
	"continue": "cherry-pickの途中でコンフリクトが解消された後、処理を再開します。\n使用例: git cherry-pick --continue",
	"abort":    "cherry-pickの処理を中止し、元の状態に戻します。\n使用例: git cherry-pick --abort",
	"n":         "コミットを適用するが、自動的にコミットはしません。\n使用例: git cherry-pick -n <commit>",
	"x":         "cherry-pickによって生成されたコミットメッセージに、元のコミットIDを追加します。\n使用例: git cherry-pick -x <commit>",
}

var cherryPickLong = `cherry-pickコマンドのヘルプを表示するコマンドです。

git cherry-pickコマンドは、指定されたコミットを現在のブランチに適用します。

基本的な使い方:
  git cherry-pick <commit>

オプション:
  --continue  コンフリクトが解消された後に処理を再開
  --abort     cherry-pickの処理を中止
  -n          コミットを適用するが、自動的にコミットはしない
  -x          コミットメッセージに元のコミットIDを追加

例:
  git cherry-pick abc123
  git cherry-pick --continue
  git cherry-pick --abort
  git cherry-pick -n abc123
  git cherry-pick -x abc123`

var cherryPickRun = `git cherry-pick:

git cherry-pickコマンドは、指定されたコミットを現在のブランチに適用します。

基本的な使い方:
  git cherry-pick <commit>

オプション:
  --continue  コンフリクトが解消された後に処理を再開
  --abort     cherry-pickの処理を中止
  -n          コミットを適用するが、自動的にコミットはしない
  -x          コミットメッセージに元のコミットIDを追加

例:
  git cherry-pick abc123
  git cherry-pick --continue
  git cherry-pick --abort
  git cherry-pick -n abc123
  git cherry-pick -x abc123`

// cherryPickCmd represents the cherry-pick command
var cherryPickCmd = &cobra.Command{
	Use:   "cherry-pick",
	Short: "cherry-pickコマンドのヘルプを表示するコマンドです。",
	Long:  cherryPickLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cherryPickRun)
			return
		}

		for _, arg := range args {
			if desc, ok := cherryPickOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cherryPickCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cherryPickCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cherryPickCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
