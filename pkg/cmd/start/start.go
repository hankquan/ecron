package start

import (
	"fmt"
	"github.com/spf13/cobra"
	"hankquan.top/ecron/pkg/store/crontab"
	"hankquan.top/ecron/pkg/util"
	"strconv"
)

type StartOptions struct {
	lineNumber int
}

func NewStartCommand() *cobra.Command {
	startOptions := &StartOptions{}

	startCmd := &cobra.Command{
		Use:   "start INDEX",
		Short: "Start the service",
		Run: func(cmd *cobra.Command, args []string) {
			// Implement the functionality here
			fmt.Println("Service started")
			util.CheckError(startOptions.Validate(args))
			util.CheckError(startOptions.Run())
		},
	}
	return startCmd
}

func (startOptions *StartOptions) Validate(args []string) error {
	lineNumber, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid index: %s", args[0])
	}
	startOptions.lineNumber = lineNumber
	return nil
}

func (startOptions *StartOptions) Run() error {
	fmt.Printf("Starting service at line number %d\n", startOptions.lineNumber)
	return crontab.StartCronEntry(startOptions.lineNumber)
}
