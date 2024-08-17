/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resetOptionDescriptions = map[string]string{
	"soft":   "--soft オプションは、インデックスと作業ツリーをそのままにしてHEADを指定したコミットに移動します。\n使用例: git reset --soft HEAD~1 (一つ前のコミットを取り消しますが、ファイルの変更はそのまま残ります)",
	"mixed":  "--mixed オプションは、インデックスを指定したコミットにリセットしますが、作業ツリーは変更しません。\n使用例: git reset --mixed HEAD~1 (一つ前のコミットを取り消し、変更はステージングされます)",
	"hard":   "--hard オプションは、インデックスと作業ツリーの両方を指定したコミットにリセットします。\n使用例: git reset --hard HEAD~1 (一つ前のコミットを取り消し、すべての変更を破棄します)",
	"merge":  "--merge オプションは、競合していない場合にインデックスと作業ツリーをリセットします。\n使用例: git reset --merge (マージ中に使用し、競合がない場合にリセットします)",
	"keep":   "--keep オプションは、競合していない場合に作業ツリーをそのままにしてインデックスをリセットします。\n使用例: git reset --keep (変更を保持しつつリセットしますが、競合がある場合はエラーを返します)",
}

var resetLong = `resetコマンドのヘルプを表示するコマンドです。

git resetコマンドは、HEADを指定した位置にリセットし、インデックスと作業ツリーの状態を変更します。オプションによって、どの範囲がリセットされるかが異なります。

基本的な使い方:
  git reset [<tree-ish>] [--soft | --mixed | --hard | --merge | --keep] [<pathspec>...]

オプション:
  --soft     HEADを移動し、インデックスと作業ツリーはそのままにします
  --mixed    インデックスをリセットし、作業ツリーは変更しません（デフォルト）
  --hard     インデックスと作業ツリーの両方をリセットします
  --merge    インデックスと作業ツリーをリセットしますが、競合がある場合はリセットしません
  --keep     インデックスをリセットし、作業ツリーをそのままにしますが、競合がある場合はリセットしません

例:
  git reset --soft HEAD~1     # 一つ前のコミットを取り消しますが、ファイルの変更はそのまま残ります
  git reset --mixed HEAD~1    # 一つ前のコミットを取り消し、変更はステージングされます
  git reset --hard HEAD~1     # 一つ前のコミットを取り消し、すべての変更を破棄します
  git reset --merge           # マージ中に使用し、競合がない場合にリセットします
  git reset --keep            # 変更を保持しつつリセットしますが、競合がある場合はエラーを返します`

var resetRun = `git reset:

git resetコマンドは、HEADを指定した位置にリセットし、インデックスと作業ツリーの状態を変更します。オプションによって、どの範囲がリセットされるかが異なります。

基本的な使い方:
  git reset [<tree-ish>] [--soft | --mixed | --hard | --merge | --keep] [<pathspec>...]

オプション:
  --soft     HEADを移動し、インデックスと作業ツリーはそのままにします
  --mixed    インデックスをリセットし、作業ツリーは変更しません（デフォルト）
  --hard     インデックスと作業ツリーの両方をリセットします
  --merge    インデックスと作業ツリーをリセットしますが、競合がある場合はリセットしません
  --keep     インデックスをリセットし、作業ツリーをそのままにしますが、競合がある場合はリセットしません

例:
  git reset --soft HEAD~1     # 一つ前のコミットを取り消しますが、ファイルの変更はそのまま残ります
  git reset --mixed HEAD~1    # 一つ前のコミットを取り消し、変更はステージングされます
  git reset --hard HEAD~1     # 一つ前のコミットを取り消し、すべての変更を破棄します
  git reset --merge           # マージ中に使用し、競合がない場合にリセットします
  git reset --keep            # 変更を保持しつつリセットしますが、競合がある場合はエラーを返します`

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "resetコマンドのヘルプを表示するコマンドです。",
	Long:  resetLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(resetRun)
			return
		}

		for _, arg := range args {
			if desc, ok := resetOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
