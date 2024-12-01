package util

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
