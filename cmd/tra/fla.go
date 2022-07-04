package tra

import (
	"time"

	"github.com/spf13/cobra"
)

type fla struct {
	Duration time.Duration
	Time     time.Time
	time     string
}

func (f *fla) Create(cmd *cobra.Command) {
	cmd.Flags().DurationVarP(&f.Duration, "dur", "d", 0, "Duration string for deriving backup data end date, e.g. 24h.")
	cmd.Flags().StringVarP(&f.time, "tim", "t", "", "Time string for backup data start date in form of yy-mm-dd.")
}

func (f *fla) Verify() {
	if f.time != "" {
		tim, err := time.Parse("06-01-02", f.time)
		if err != nil {
			panic(err)
		}

		f.Time = tim
	}
}
