package com

import (
	"os"

	"github.com/spf13/cobra"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		err := cmd.Root().GenBashCompletion(os.Stdout)
		if err != nil {
			panic(err)
		}
	case "zsh":
		err := cmd.Root().GenZshCompletion(os.Stdout)
		if err != nil {
			panic(err)
		}
	case "fish":
		err := cmd.Root().GenFishCompletion(os.Stdout, true)
		if err != nil {
			panic(err)
		}
	case "powershell":
		err := cmd.Root().GenPowerShellCompletion(os.Stdout)
		if err != nil {
			panic(err)
		}
	}
}
