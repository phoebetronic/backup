package red

import (
	"github.com/phoebetron/backup/cmd/red/upl"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

const (
	use = "red"
	sho = "Manage redis resources."
	lon = "Manage redis resources."
)

type Config struct{}

func New(con Config) (*cobra.Command, error) {
	var err error

	// --------------------------------------------------------------------- //

	var cmdUpl *cobra.Command
	{
		c := upl.Config{}

		cmdUpl, err = upl.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// --------------------------------------------------------------------- //

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			RunE:  (&run{}).run,
		}
	}

	{
		c.AddCommand(cmdUpl)
	}

	return c, nil
}
