package remove

import (
	"fmt"
	"github.com/spf13/cobra"
	"hankquan.top/ecron/pkg/store/crontab"
	"hankquan.top/ecron/pkg/util"
	"strconv"
)

type RemoveOptions struct {
	lineNumber int
}

func NewRemoveCommand() *cobra.Command {
	removeOptions := &RemoveOptions{}

	var removeCmd = &cobra.Command{
		Use:   "remove INDEX",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("AddOptions before validate: %+v\n", removeOptions)
			util.CheckError(removeOptions.Validate(args))
			//generate backup file
			fmt.Printf("AddOptions before run: %+v\n", removeOptions)
			util.CheckError(removeOptions.Run())
			//add change log to history file if succeed
		},
	}

	return removeCmd
}

func (removeOptions *RemoveOptions) Validate(args []string) error {
	lineNumber, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid index: %s", args[0])
	}
	removeOptions.lineNumber = lineNumber
	return nil
}

func (removeOptions *RemoveOptions) Run() error {
	return crontab.DeleteCronEntry(removeOptions.lineNumber)
}
