package win

import (
	"github.com/spf13/cobra"
)

const (
	use = "win"
	sho = "Generate window frames of training input data."
	lon = `
Generate window frames of training input data.

    backup win --tim 22-06-01 --csv

    backup win --tim 22-06-01 --dra
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
