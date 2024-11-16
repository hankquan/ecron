/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"github.com/spf13/cobra"
	"hankquan.top/ecron/pkg/cmd"
	"hankquan.top/ecron/pkg/crontab"
	"hankquan.top/ecron/pkg/util"
)

type ListOption struct {
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all cron jobs",
	Long:  `List all cron jobs by crontab -l`,
	Run: func(cmd *cobra.Command, args []string) {
		cronLines := crontab.GetCronElements()
		util.PrintTable(cronLines)
	},
}

func init() {
	cmd.RootCmd.AddCommand(listCmd)
	//cmd.rootCmd.AddCommand(listCmd)
}
