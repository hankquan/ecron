/*
Copyright © 2024 Hank Quan <hankwayup@outlook.com>
*/
package main

import (
	"fmt"
	"hankquan.top/ecron/pkg/cmd"
)

var (
	Version   string
	BuildTime string
)

func main() {
	command := cmd.NewEcronCommand(Version, BuildTime)
	err := command.Execute()
	if err != nil {
		fmt.Print("Error: ", err.Error())
	}
}
