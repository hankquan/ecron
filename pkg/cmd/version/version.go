package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewVersionCommand(version string, buildTime string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long:  `Show version and build time`,
		Run: func(command *cobra.Command, args []string) {
			fmt.Println("Version:   ", version)
			fmt.Println("BuildTime: ", buildTime)
		},
	}
}
