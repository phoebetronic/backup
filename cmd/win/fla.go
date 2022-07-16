package win

import (
	"time"

	"github.com/spf13/cobra"
)

type fla struct {
	CSV bool
	Dra bool
	Ind int
	Tim time.Time
	tim string
}

func (f *fla) Create(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&f.CSV, "csv", "", false, "Whether to write all generated training data to a .csv file.")
	cmd.Flags().BoolVarP(&f.Dra, "dra", "", false, "Whether to plot some graphs based on the generated training data.")
	cmd.Flags().IntVarP(&f.Ind, "ind", "", 0, "The window frame start index to plot graphs for.")
	cmd.Flags().StringVarP(&f.tim, "tim", "", "", "Time string for backup data start date in form of yy-mm-dd.")
}

func (f *fla) Verify() {
	if !f.CSV && !f.Dra {
		panic("either of --csv/--dra must be given")
	}

	if f.tim == "" {
		panic("-t/--tim must not be empty")
	}

	{
		tim, err := time.Parse("06-01-02", f.tim)
		if err != nil {
			panic(err)
		}

		f.Tim = tim
	}

	if f.Tim.Day() != 1 {
		panic("-t/--tim must specify the first day of the month")
	}
}
