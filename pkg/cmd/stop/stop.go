package stop

import (
	"fmt"
	"github.com/spf13/cobra"
	"hankquan.top/ecron/pkg/store/crontab"
	"hankquan.top/ecron/pkg/util"
	"strconv"
)

type StopOptions struct {
	lineNumber int
}

func NewStopCommand() *cobra.Command {
	stopOptions := &StopOptions{}
	StopCmd := &cobra.Command{
		Use:   "stop INDEX",
		Short: "Stop the service",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckError(stopOptions.Validate(args))
			util.CheckError(stopOptions.Run())
		},
	}
	return StopCmd
}

func (StopOptions *StopOptions) Validate(args []string) error {
	lineNumber, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid index: %s", args[0])
	}
	StopOptions.lineNumber = lineNumber
	return nil
}

func (StopOptions *StopOptions) Run() error {
	fmt.Printf("Stoping service at line number %d\n", StopOptions.lineNumber)
	return crontab.StopCronEntry(StopOptions.lineNumber)
}
