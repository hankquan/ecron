/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"
	"hankquan.top/ecron/pkg/cmd"

	"github.com/spf13/cobra"
)

type AddOptions struct {
	minutely bool
	hourly   bool
	daily    bool
	weekly   bool
	min      int    //0~60 <0 -->0 || >59 -->0
	at       string //1am,7pm,14pm
	on       string //monday, Monday, Mon, mon
	expr     string
	prompt   string
}

func NewAddCommand() *cobra.Command {
	addOptions := &AddOptions{}

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Add called")
			//validate
			for _, arg := range args {
				fmt.Println("arg-->", arg)
			}
			//run
			err := addOptions.Run()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Done!")
		},
	}
	addCmd.Flags().BoolVar(&addOptions.minutely, "minutely", false, "run every minute")
	return addCmd
}

func init() {

	addCommand := NewAddCommand()
	cmd.RootCmd.AddCommand(addCommand)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this addCommand
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this addCommand
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (addOptions *AddOptions) Run() error {
	//check

	return nil
}
