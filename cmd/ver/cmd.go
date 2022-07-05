package ver

import (
	"github.com/spf13/cobra"
)

const (
	use = "ver"
	sho = "Print version information of this command line tool."
	lon = "Print version information of this command line tool."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
		}
	}

	return c, nil
}
