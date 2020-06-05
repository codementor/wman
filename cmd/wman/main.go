package main

import (
	"os"

	"github.com/codementor/wman/pkg/cmd"
)

func main() {
	if err := cmd.NewWmanCmd().Execute(); err != nil {
		os.Exit(-1)
	}
}
