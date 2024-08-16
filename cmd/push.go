/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushOptionDescriptions = map[string]string{
	"u":      "\n-uオプションは、ローカルの現在のブランチを指定したリモートブランチと紐付けるために使用されます。\nこのオプションを指定して git push を行うと、以降そのブランチでの git push や git pull コマンドは、リモートブランチ名を省略して実行できるようになります。\nつまり、2回目以降のプッシュでは、git push だけでリモートの対応するブランチにプッシュされるようになります。\n\n\n例えば、feature/new-feature という名前のローカルブランチがあり、それを origin リモートの同名のブランチにプッシュしたい場合、以下のコマンドを使います。\n\n  $ git push -u origin feature/new-feature\n\nこれにより、feature-branch ブランチは origin/feature-branch に紐付けられます。\nその後は、単に git push と入力するだけで origin/feature-branch にプッシュできるようになります。",
	"set-upstream":      "\n-uオプションは、ローカルの現在のブランチを指定したリモートブランチと紐付けるために使用されます。\nこのオプションを指定して git push を行うと、以降そのブランチでの git push や git pull コマンドは、リモートブランチ名を省略して実行できるようになります。\nつまり、2回目以降のプッシュでは、git push だけでリモートの対応するブランチにプッシュされるようになります。\n\n\n例えば、feature/new-feature という名前のローカルブランチがあり、それを origin リモートの同名のブランチにプッシュしたい場合、以下のコマンドを使います。\n\n  $ git push -u origin feature/new-feature\n\nこれにより、feature-branch ブランチは origin/feature-branch に紐付けられます。\nその後は、単に git push と入力するだけで origin/feature-branch にプッシュできるようになります。",
	"force":  "\n--forceオプションは、ローカルの変更を強制的にリモートリポジトリにプッシュするために使用されます。\nこのオプションを使用すると、リモートリポジトリの履歴が上書きされる可能性があるため、慎重に扱う必要があります。\n通常、リモートに存在するブランチとローカルのブランチに異なるコミットがある場合、Gitはデータの損失を防ぐためにプッシュを拒否します。\nしかし、--forceオプションを使用すると、この保護を無視して強制的にローカルの履歴をリモートに適用します。\n\n\n例えば、mainブランチに強制的にプッシュしたい場合は、以下のコマンドを使います。\n\n  $ git push --force origin main\n\nこれは、リモートの main ブランチにあるすべてのコミットを上書きし、ローカルの main ブランチの状態を反映させます。\n注意点として、--force を使うとリモートにある他の人の作業内容が失われる可能性があるため、特にチームでの作業中には使用する際に十分な確認が必要です。",
	"tags":   "\n--tagsオプションは、リモートリポジトリにタグを含めてすべてのリファレンスをプッシュするために使用されます。\nGitのタグは、特定のコミットに対して「名前」をつける機能で、通常はリリースのバージョン管理などに利用されます。\n通常の git push コマンドでは、ブランチの更新のみがプッシュされ、タグはプッシュされません。\nしかし、--tags オプションを指定することで、ローカルリポジトリに存在するすべてのタグがリモートリポジトリに送信されます。\n\n\n例えば、以下のコマンドを実行することで、すべてのタグをリモートの origin にプッシュすることができます。\n\n  $ git push --tags origin\n\nこのコマンドにより、ローカルに存在するすべてのタグがリモートリポジトリに反映されます。\nタグは通常、リリースバージョンの固定や重要なコミットポイントの識別に使用されるため、これをリモートに共有することで、他の開発者と正確な状態を共有することができます。",
}

var pushLong = `pushコマンドのヘルプを表示するコマンドです。

git pushコマンドは、ローカルリポジトリの変更をリモートリポジトリに送信します。

基本的な使い方:
  git push <リモート> <ブランチ>

オプション:
  -u, --set-upstream    アップストリームとして設定
  --force               強制的にプッシュ
  --tags                タグをプッシュ

例:
  git push origin main
  git push -u origin feature/new-feature
  git push --force origin main
  git push --tags origin`

var pushRun = `git push:

git pushコマンドは、ローカルリポジトリの変更をリモートリポジトリに送信します。

基本的な使い方:
  git push <リモート> <ブランチ>

オプション:
  -u, --set-upstream    アップストリームとして設定
  --force               強制的にプッシュ
  --tags                タグをプッシュ

例:
  git push origin main
  git push -u origin feature/new-feature
  git push --force origin main
  git push --tags origin`

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "pushコマンドのヘルプを表示するコマンドです。",
	Long:  pushLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(pushRun)
			return
		}

		for _, arg := range args {
			if desc, ok := pushOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
