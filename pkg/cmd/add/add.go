package add

import (
	"fmt"
	"github.com/spf13/cobra"
	"hankquan.top/ecron/pkg/store/crontab"
	"hankquan.top/ecron/pkg/util"
)

func NewAddCommand() *cobra.Command {
	cronOptions := &CronOptions{}

	var addCmd = &cobra.Command{
		Use:   "add [--hourly] [--daily] [--weekly] [--minutely] [--expr] [--prompt] [--past] [--at] [--on] command",
		Short: "Add a new job to crontab",
		Long:  `Add a new job to crontab by specifying the command and schedule.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args[0]) <= 0 {
				fmt.Println("Invalid command")
				return
			}
			cronOptions.cmd = args[0]
			util.CheckError(cronOptions.RunAdd())
		},
	}

	addCmd.Flags().BoolVar(&cronOptions.minutely, "minutely", false, "run every minute")
	addCmd.Flags().BoolVar(&cronOptions.hourly, "hourly", false, "run every hour")
	addCmd.Flags().BoolVar(&cronOptions.daily, "daily", false, "run every day")
	addCmd.Flags().BoolVar(&cronOptions.weekly, "weekly", false, "run every week")
	addCmd.Flags().StringVar(&cronOptions.at, "at", "", "specify an hour for --daily")
	addCmd.Flags().StringArrayVar(&cronOptions.on, "on", nil, "specify weekdays for --weekly")

	addCmd.Flags().StringVar(&cronOptions.expr, "expr", "", "run by cron expression")
	addCmd.Flags().StringVar(&cronOptions.prompt, "", "", "auto-generate cron expression by prompt")
	return addCmd
}

func (cronOptions *CronOptions) RunAdd() error {
	cronExpr, err := cronOptions.GenerateCronExpression()
	if err != nil {
		return err
	}
	fmt.Print("generate cron expression: ", cronExpr)
	return crontab.AddCronEntry(cronExpr, cronOptions.cmd)
}
