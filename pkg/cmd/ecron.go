package cmd

import (
	"github.com/spf13/cobra"
	"hankquan.top/ecron/pkg/cmd/add"
	"hankquan.top/ecron/pkg/cmd/list"
	"hankquan.top/ecron/pkg/cmd/remove"
	"hankquan.top/ecron/pkg/cmd/start"
	"hankquan.top/ecron/pkg/cmd/stop"
	"hankquan.top/ecron/pkg/cmd/version"
)

func NewEcronCommand(buildVersion string, buildTime string) *cobra.Command {
	var ecronCommand = &cobra.Command{
		Use:   "ecron [command] [flags]",
		Short: "ecron is an easy cli tool for managing crontab",
		Long: `ecron is a simple and intuitive command-line tool designed for managing Linux crontab scheduled tasks.
Find more information at: https://github.com/hankquan/ecron`,
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	ecronCommand.SetHelpCommand(&cobra.Command{
		Use:    "help",
		Hidden: true,
	})

	ecronCommand.AddCommand(list.NewListCommand())
	ecronCommand.AddCommand(add.NewAddCommand())
	ecronCommand.AddCommand(add.NewEditCommand())
	ecronCommand.AddCommand(remove.NewRemoveCommand())
	ecronCommand.AddCommand(version.NewVersionCommand(buildVersion, buildTime))
	ecronCommand.AddCommand(start.NewStartCommand())
	ecronCommand.AddCommand(stop.NewStopCommand())
	return ecronCommand
}
