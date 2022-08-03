package ens

import (
	"github.com/spf13/cobra"
)

const (
	use = "ens"
	sho = "Generate ensemble training data."
	lon = `Generate ensemble training data. In order to write the generated ensemble
training data into CSV files use the command below.

    backup ens --tim 22-06-01
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
