package win

import (
	"github.com/spf13/cobra"
)

const (
	use = "win"
	sho = "Generate window frames of training input data."
	lon = `Generate window frames of training input data. In order to write the generated
training data into a .csv file use the command below.

    backup win --tim 22-06-01 --csv

Further, graphs of the generated training data can be plotted by selecting a
window frame index determining which window frames to plot. Note that 10 graphs
will be plotted at a time, which will all be very similar due to the sequential
nature of the generated window frames.

    backup win --tim 22-06-01 --dra --ind 20
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
