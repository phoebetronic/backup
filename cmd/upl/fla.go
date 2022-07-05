package upl

import (
	"time"

	"github.com/spf13/cobra"
)

type fla struct {
	Time time.Time
	time string
}

func (f *fla) Create(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.time, "tim", "t", "", "Time string for backup data start date in form of yy-mm-dd.")
}

func (f *fla) Verify() {
	if f.time == "" {
		panic("-t/--tim must not be empty")
	}

	{
		tim, err := time.Parse("06-01-02", f.time)
		if err != nil {
			panic(err)
		}

		f.Time = tim
	}

	if f.Time.Day() != 1 {
		panic("-t/--tim must specify the first day of the month")
	}
}
