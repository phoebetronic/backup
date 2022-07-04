package red

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) error {
	{
		err := cmd.Help()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
