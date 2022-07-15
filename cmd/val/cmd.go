package val

import (
	"github.com/spf13/cobra"
)

const (
	use = "val"
	sho = "Validate raw backup data."
	lon = `Validate raw backup data. Before backups can be inspected and validated, the
monthly partitions have to be downloaded from S3. Below is shown how to read
backup trades of a specific month back from Redis once they have been
downloaded, in order to check the content of the downloaded partitions.

    backup val --tim 22-06-01
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
