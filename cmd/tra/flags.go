package tra

import (
	"time"

	"github.com/spf13/cobra"
)

type flags struct {
	Exchange string
	Asset    string
	Duration time.Duration
	Time     time.Time
	time     string
}

func (f *flags) Create(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Exchange, "exc", "e", "", "The exchange to backup trades from, e.g. ftx.")
	cmd.Flags().StringVarP(&f.Asset, "ass", "a", "", "The asset to backup trades for, e.g. eth.")
	cmd.Flags().DurationVarP(&f.Duration, "dur", "d", 0, "Duration string for deriving backup data end date, e.g. 24h.")
	cmd.Flags().StringVarP(&f.time, "tim", "t", "", "Time string for backup data start date in form of yy-mm-dd.")
}

func (f *flags) Verify() {
	if f.Exchange == "" {
		panic("-e/--exc must not be empty")
	}

	if f.Asset == "" {
		panic("-a/--ass must not be empty")
	}

	if f.time != "" {
		tim, err := time.Parse("06-01-02", f.time)
		if err != nil {
			panic(err)
		}

		f.Time = tim.UTC()
	}
}
