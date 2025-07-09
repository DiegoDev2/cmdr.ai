package main

import (
	"os"

	"cmdr.ai/cmd"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		cmd.InitShellIntegration()
		os.Exit(0)
	}
	cmd.Main()
}
