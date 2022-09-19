package ord

import (
	"github.com/spf13/cobra"
)

const (
	use = "ord"
	sho = "Backup raw orderbook data."
	lon = `Backup raw orderbook data. Backing up raw orderbook data can only
happen on time and not retrospectively. This is because there is no historical
archive for orderbook data. Therefore orderbook backups have to be created for a
specific timespan without interruption. For now only dYdX orderbook backups are
supported.

    backup ord --exc dydx --ass eth
`
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var f *flags
	{
		f = &flags{}
	}

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{flags: f}).run,
		}
	}

	{
		f.Create(c)
	}

	return c, nil
}
