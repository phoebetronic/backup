package upl

import (
	"github.com/spf13/cobra"
)

const (
	use = "upl"
	sho = "Upload backups to S3."
	lon = `Upload backups to S3. When uploading backups to S3, trades are partitioned into
monthly packages of single ticks. Below is shown how to index and upload a
specific month of trades.

    backup upl --kin tra --exc ftx --ass eth --tim 22-06-01

Orders from orderbook backups are partitioned in hours. The command below shows
how to upload a particular hour of orderbook backups.

    backup upl --kin ord --exc dydx --ass eth --tim 22-09-19T14:00:00
`
)

type Config struct{}

func New(con Config) (*cobra.Command, error) {
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
