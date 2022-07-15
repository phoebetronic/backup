package dow

import (
	"github.com/spf13/cobra"
)

const (
	use = "dow"
	sho = "Download backups from S3."
	lon = `Download backups from S3. When downloading backups from S3, trades are
downloaded in monhtly partitions of single ticks. Below is shown how to feed
backup trades of a specific month back into Redis.

    backup dow --tim 22-06-01
`
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
