/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reflogOptionDescriptions = map[string]string{
	"expire": "\nreflog expireオプションは、リファレンスログを手動で削除するために使用されます。\nGitは自動的に古いリファレンスログを削除しますが、このオプションを使って特定の期間より古いログを手動で削除することができます。\n\n例えば、30日以上前のログを削除したい場合は、以下のコマンドを使用します。\n\n  $ git reflog expire --expire=30.days.ago\n",
	"delete": "\nreflog deleteオプションは、特定のエントリをリファレンスログから削除するために使用されます。\n通常はすべてのログを保持しますが、このオプションを使って特定の不要なエントリを削除することができます。\n\n例えば、最新のログエントリを削除するには、以下のコマンドを使用します。\n\n  $ git reflog delete HEAD@{0}\n",
	"show":   "\nreflog showオプションは、リファレンスログのエントリを表示するために使用されます。\nこのオプションを使うと、特定のブランチやリファレンスの履歴を確認することができます。\n\n例えば、HEADのリファレンスログを表示するには、以下のコマンドを使用します。\n\n  $ git reflog show HEAD\n",
}

var reflogLong = `reflogコマンドのヘルプを表示するコマンドです。

git reflogコマンドは、ローカルリポジトリで行ったコミットやチェックアウトなどの履歴を表示します。

基本的な使い方:
  git reflog <オプション> [リファレンス]

オプション:
  expire    リファレンスログを手動で削除
  delete    特定のリファレンスログエントリを削除
  show      リファレンスログを表示

例:
  git reflog
  git reflog show HEAD
  git reflog expire --expire=30.days.ago
  git reflog delete HEAD@{0}`

var reflogRun = `git reflog:

git reflogコマンドは、ローカルリポジトリで行ったすべてのコミット、リセット、チェックアウトなどの操作履歴を追跡します。

基本的な使い方:
  git reflog <オプション> [リファレンス]

オプション:
  expire    リファレンスログを手動で削除
  delete    特定のリファレンスログエントリを削除
  show      リファレンスログを表示

例:
  git reflog
  git reflog show HEAD
  git reflog expire --expire=30.days.ago
  git reflog delete HEAD@{0}`

// reflogCmd represents the reflog command
var reflogCmd = &cobra.Command{
	Use:   "reflog",
	Short: "reflogコマンドのヘルプを表示するコマンドです。",
	Long:  reflogLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(reflogRun)
			return
		}

		for _, arg := range args {
			if desc, ok := reflogOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(reflogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reflogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reflogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
