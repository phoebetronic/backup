package com

import (
	"github.com/spf13/cobra"
)

const (
	use = "com"
	sho = "Generate shell completions."
	lon = `Supported positional arguments and respective shell completions are
as follows.

    bash
    fish
    powershell
    zsh

Generating zsh completion for Oh My Zsh can be done by writing the
generated completion to the custom plugin folder.

    mkdir -p ~/.oh-my-zsh/custom/plugins/backup && backup com zsh > ~/.oh-my-zsh/custom/plugins/backup/_backup

	`
)

type Config struct{}

func New(con Config) (*cobra.Command, error) {
	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:                   use,
			Short:                 sho,
			Long:                  lon,
			DisableFlagsInUseLine: true,
			ValidArgs:             []string{"bash", "fish", "powershell", "zsh"},
			Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
			Run:                   (&run{}).run,
		}
	}

	return c, nil
}
