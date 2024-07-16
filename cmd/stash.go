/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stashOptionDescriptions = map[string]string{
	"save":       "save オプションは、現在の作業ディレクトリの変更をスタッシュに保存します。\n使用例: git stash save \"作業内容\"",
	"pop":        "pop オプションは、最も最近のスタッシュを適用し、そのエントリを削除します。\n使用例: git stash pop",
	"list":       "list オプションは、すべてのスタッシュエントリの一覧を表示します。\n使用例: git stash list",
	"apply":      "apply オプションは、指定したスタッシュエントリを適用しますが、スタッシュエントリは削除されません。\n使用例: git stash apply stash@{0}",
	"drop":       "drop オプションは、指定したスタッシュエントリを削除します。\n使用例: git stash drop stash@{0}",
	"clear":      "clear オプションは、すべてのスタッシュエントリを削除します。\n使用例: git stash clear",
}

var stashLong = `stashコマンドのヘルプを表示するコマンドです。

git stashコマンドは、作業中の変更を一時的に保存し、作業ディレクトリをクリーンな状態に戻します。

基本的な使い方:
  git stash <オプション> <引数>

オプション:
  save        現在の作業ディレクトリの変更をスタッシュに保存
  pop         最も最近のスタッシュを適用し、そのエントリを削除
  list        すべてのスタッシュエントリの一覧を表示
  apply       指定したスタッシュエントリを適用
  drop        指定したスタッシュエントリを削除
  clear       すべてのスタッシュエントリを削除

例:
  git stash save "作業内容"
  git stash pop
  git stash list
  git stash apply stash@{0}
  git stash drop stash@{0}
  git stash clear`

var stashRun = `git stash:

git stashコマンドは、作業中の変更を一時的に保存し、作業ディレクトリをクリーンな状態に戻します。

基本的な使い方:
  git stash <オプション> <引数>

オプション:
  save        現在の作業ディレクトリの変更をスタッシュに保存
  pop         最も最近のスタッシュを適用し、そのエントリを削除
  list        すべてのスタッシュエントリの一覧を表示
  apply       指定したスタッシュエントリを適用
  drop        指定したスタッシュエントリを削除
  clear       すべてのスタッシュエントリを削除

例:
  git stash save "作業内容"
  git stash pop
  git stash list
  git stash apply stash@{0}
  git stash drop stash@{0}
  git stash clear`

// stashCmd represents the stash command
var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "stashコマンドのヘルプを表示するコマンドです。",
	Long:  stashLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(stashRun)
			return
		}

		for _, arg := range args {
			if desc, ok := stashOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}


func init() {
	rootCmd.AddCommand(stashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
