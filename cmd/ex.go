/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var exOptionDescriptions = map[string]string{
	"1": "コミットの取消\n\n使用例:\n\n// 直前のコミットを取り消す\n  git reset --soft HEAD~1",
	"2": "直前のコミットに変更を追加する\n\n使用例:\n\n// 追加したい変更がステージングされている状態で \n  git commit --amend --no-edit",
	"3": "直前のコミットメッセージを修正する\n\n使用例:\n  git commit --amend",
	"4": "ブランチを間違えた\n\n使用例:\n  git stash\n  git checkout [正しいブランチ名]\n  git stash pop\n\n  これにより差分を別ブランチに移動することができます。",
	"5": "リモートブランチを削除する\n\n使用例:\n  git push origin --delete [リモートブランチ名]",
	"6": "マージ時のコンフリクトを手動で解決した後にマージする\n\n使用例:\n  git add <ファイル名>\ngit merge --continue",
}

var exLong = `exコマンドの使用例を表示するコマンドです。`

var exRun = `使用例：
1. コミットの取消
2. 直前のコミットに変更を追加する
3. 直前のコミットメッセージを修正する
4. コミットするブランチを間違えた
5. リモートブランチを削除する
6. マージ時のコンフリクトを手動で解決した後にマージする`

// exCmd represents the ex command
var exCmd = &cobra.Command{
	Use:   "ex",
	Short: "gitコマンドの使用例を表示するコマンドです。",
	Long: exLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(exRun)
			return
		}

		for _, arg := range args {
			if desc, ok := exOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(exCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
