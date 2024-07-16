/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pullOptionDescriptions = map[string]string{
	"rebase": "--rebase オプションは、リモートブランチの変更をマージする代わりにリベースを行います。\n使用例: git pull --rebase",
	"all":    "--all オプションは、すべてのリモートリポジトリから変更を取得してマージします。\n使用例: git pull --all",
}

var pullLong = `pullコマンドのヘルプを表示するコマンドです。

git pullコマンドは、リモートリポジトリから変更を取得し、ローカルリポジトリにマージします。

基本的な使い方:
  git pull <リモート> <ブランチ>

オプション:
  --rebase              リモートブランチの変更をマージする代わりにリベースを行います
  --all                 すべてのリモートリポジトリから変更を取得してマージします

例:
  git pull origin main
  git pull --rebase origin main
  git pull --all`

var pullRun = `git pull:

git pullコマンドは、リモートリポジトリから変更を取得し、ローカルリポジトリにマージします。

基本的な使い方:
  git pull <リモート> <ブランチ>

オプション:
  --rebase              リモートブランチの変更をマージする代わりにリベースを行います
  --all                 すべてのリモートリポジトリから変更を取得してマージします

例:
  git pull origin main
  git pull --rebase origin main
  git pull --all`

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pullコマンドのヘルプを表示するコマンドです。",
	Long:  pullLong,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(pullRun)
			return
		}

		for _, arg := range args {
			if desc, ok := pullOptionDescriptions[arg]; ok {
				fmt.Printf("%s\n\n", desc)
			} else {
				fmt.Printf("不明なオプション: %s\n\n", arg)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
