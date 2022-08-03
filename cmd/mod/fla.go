package mod

import (
	"time"

	"github.com/spf13/cobra"
)

type fla struct {
	Buc float32
	Ind int
	Tim time.Time
	tim string
}

func (f *fla) Create(cmd *cobra.Command) {
	cmd.Flags().Float32VarP(&f.Buc, "buc", "", 0, "The target bucket to generate datasets for.")
	cmd.Flags().IntVarP(&f.Ind, "ind", "", 0, "The window frame start index to plot graphs for.")
	cmd.Flags().StringVarP(&f.tim, "tim", "", "", "Time string for backup data start date in form of yy-mm-dd.")
}

func (f *fla) Verify() {
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
