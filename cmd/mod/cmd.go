package mod

import (
	"github.com/spf13/cobra"
)

const (
	use = "mod"
	sho = "Generate model training data."
	lon = `Generate model training data. In order to write the generated model training
data into CSV files use the command below.

    backup mod --tim 22-06-01
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
