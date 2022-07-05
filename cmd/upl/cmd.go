package upl

import (
	"github.com/spf13/cobra"
)

const (
	use = "upl"
	sho = "Upload backups to S3."
	lon = "Upload backups to S3."
)

type Config struct{}

func New(con Config) (*cobra.Command, error) {
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
			RunE:  (&run{fla: f}).run,
		}
	}

	{
		f.Create(c)
	}

	return c, nil
}
