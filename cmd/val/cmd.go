package val

import (
	"github.com/spf13/cobra"
)

const (
	use = "val"
	sho = "Validate raw backup data"
	lon = "Validate raw backup data"
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var f *fla
	{
		f = &fla{}
	}

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{cmdfla: f}).run,
		}
	}

	{
		f.Create(c)
	}

	return c, nil
}
