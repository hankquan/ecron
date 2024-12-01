package add

import (
	"fmt"
	"github.com/spf13/cobra"
	"hankquan.top/ecron/pkg/store/crontab"
	"hankquan.top/ecron/pkg/util"
	"strconv"
)

func NewEditCommand() *cobra.Command {
	cronOptions := &CronOptions{}

	var editCmd = &cobra.Command{
		Use:   "edit INDEX [--cmd new-command] [--hourly] [--daily] [--weekly] [--minutely] [--expr] [--prompt] [--past] [--at] [--on]",
		Short: "Edit an existing crontab job by INDEX",
		Long: `Edit an existing crontab job by specifying the INDEX and new command or schedule.
For example:
ecron edit 1 --hourly --cmd /path/to/new/command # edit both command and schedule
ecron edit 2 --daily --at 12 # edit schedule only
ecron edit 3 --expr "0 0 * * * /path/to/command" # edit by cron expression
`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lineNumber, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid line number:", err)
				return
			}
			util.CheckError(cronOptions.RunEdit(lineNumber))
		},
	}

	editCmd.Flags().StringVar(&cronOptions.cmd, "cmd", "", "specify the new command for the crontab job")

	editCmd.Flags().BoolVar(&cronOptions.minutely, "minutely", false, "run every minute")
	editCmd.Flags().BoolVar(&cronOptions.hourly, "hourly", false, "run every hour")
	editCmd.Flags().BoolVar(&cronOptions.daily, "daily", false, "run every day")
	editCmd.Flags().BoolVar(&cronOptions.weekly, "weekly", false, "run every week")
	editCmd.Flags().StringVar(&cronOptions.at, "at", "", "specify an hour for --daily")
	editCmd.Flags().StringArrayVar(&cronOptions.on, "on", nil, "specify weekdays for --weekly")

	editCmd.Flags().StringVar(&cronOptions.expr, "expr", "", "run by cron expression")
	editCmd.Flags().StringVar(&cronOptions.prompt, "", "", "auto-generate cron expression by prompt")
	return editCmd
}

func (cronOptions *CronOptions) RunEdit(lineNumber int) error {
	cronExpr, err := cronOptions.GenerateCronExpression()
	if err != nil {
		return err
	}
	return crontab.EditCronEntry(lineNumber, cronExpr, cronOptions.cmd)
}
