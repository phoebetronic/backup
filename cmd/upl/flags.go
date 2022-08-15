package upl

import (
	"time"

	"github.com/spf13/cobra"
)

type flags struct {
	Exchange string
	Time     time.Time
	time     string
}

func (f *flags) Create(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Exchange, "exc", "e", "", "The exchange to upload trades for, e.g. ftx.")
	cmd.Flags().StringVarP(&f.time, "tim", "t", "", "Time string for backup data start date in form of yy-mm-dd.")
}

func (f *flags) Verify() {
	if f.Exchange == "" {
		panic("-e/--exc must not be empty")
	}

	if f.time == "" {
		panic("-t/--tim must not be empty")
	}

	{
		tim, err := time.Parse("06-01-02", f.time)
		if err != nil {
			panic(err)
		}

		f.Time = tim.UTC()
	}

	if f.Time.Day() != 1 {
		panic("-t/--tim must specify the first day of the month")
	}
}
