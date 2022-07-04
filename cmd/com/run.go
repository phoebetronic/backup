package com

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) error {
	{
		switch args[0] {
		case "bash":
			err := cmd.Root().GenBashCompletion(os.Stdout)
			if err != nil {
				return tracer.Mask(err)
			}
		case "zsh":
			err := cmd.Root().GenZshCompletion(os.Stdout)
			if err != nil {
				return tracer.Mask(err)
			}
		case "fish":
			err := cmd.Root().GenFishCompletion(os.Stdout, true)
			if err != nil {
				return tracer.Mask(err)
			}
		case "powershell":
			err := cmd.Root().GenPowerShellCompletion(os.Stdout)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
