package tra

import (
	"github.com/spf13/cobra"
)

const (
	use = "tra"
	sho = "Backup raw trades data."
	lon = `Backup raw trades data. When starting to backup trades the first time, then the
starting time of the backup period has to be provided like shown below.

    backup tra --exc ftx --dur 1h --tim 20-01-01

Further backups continue based on the latest trade known in Redis. For instance,
in order to backup another hour, starting from where we left off last time, we
can simply execute the command below.

    backup tra --exc ftx --dur 1h
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
