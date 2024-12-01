/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"github.com/spf13/cobra"
	"hankquan.top/ecron/pkg/store/crontab"
	"hankquan.top/ecron/pkg/util"
)

type ListOption struct {
}

func NewListCommand() *cobra.Command {
	listOptions := &ListOption{}
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all cron jobs",
		Long:  `List all cron jobs by crontab -l`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckError(listOptions.Run())
		},
	}
	return listCmd
}

func (ListOption *ListOption) Run() error {
	cronLines, err := crontab.GetCronEntries()
	if err != nil {
		return err
	}
	util.PrintTable(cronLines)
	return nil
}
