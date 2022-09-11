package fil

import (
	"github.com/spf13/cobra"
)

const (
	use = "fil"
	sho = "Write backup trades within a given time range from S3 to local files."
	lon = `Write backup trades within a given time range from S3 to local files. When
downloading backups from S3, trades are downloaded in monthly partitions of
single ticks. The given time range defines teh trades to write to the given file
path. Below is shown how to write backup trades within a given time range back
into local files.

    backup fil --exc ftx --ass eth --sta 22-06-01T01:00:00  --end 22-06-01T07:00:00 --dur 1m --pat /Users/xh3b4sd/phoebetron/001.json
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
