package fil

import (
	"time"

	"github.com/spf13/cobra"
)

type flags struct {
	Exc string
	Ass string
	Sta time.Time
	End time.Time
	Dur time.Duration
	Kin string
	Pat string
	sta string
	end string
}

func (f *flags) Create(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Exc, "exc", "", "", "The exchange to download trades from, e.g. ftx.")
	cmd.Flags().StringVarP(&f.Ass, "ass", "", "", "The asset to download trades for, e.g. eth.")
	cmd.Flags().DurationVarP(&f.Dur, "dur", "", 0, "The frame duration of buffered trades, e.g. 1m.")
	cmd.Flags().StringVarP(&f.Kin, "kin", "k", "", "The kind of backup data to validate, e.g. ord, tra.")
	cmd.Flags().StringVarP(&f.Pat, "pat", "", "", "The file path to write trades to, e.g. /Users/xh3b4sd/phoebetron/001.json.")
	cmd.Flags().StringVarP(&f.sta, "sta", "", "", "Time string for backup data start date in form of yy-mm-ddThh:00:00.")
	cmd.Flags().StringVarP(&f.end, "end", "", "", "Time string for backup data end date in form of yy-mm-ddThh:00:00.")
}

func (f *flags) Verify() {
	{
		if f.Exc == "" {
			panic("--exc must not be empty")
		}

		if f.Ass == "" {
			panic("--ass must not be empty")
		}

		if f.Kin == "" {
			panic("-k/--kin must not be empty")
		}
		if f.Kin != "ord" && f.Kin != "tra" {
			panic("-k/--kin must be either ord or tra")
		}
	}

	{
		if f.sta == "" {
			panic("--sta must not be empty")
		}

		sta, err := time.Parse("06-01-02T15:04:05", f.sta)
		if err != nil {
			panic(err)
		}

		f.Sta = sta.UTC()

		if f.Sta.Minute() != 0 {
			panic("--sta must not specify minutes, only full hours are supported")
		}

		if f.Sta.Second() != 0 {
			panic("--sta must not specify seconds, only full hours are supported")
		}
	}

	{
		if f.end == "" {
			panic("--end must not be empty")
		}

		end, err := time.Parse("06-01-02T15:04:05", f.end)
		if err != nil {
			panic(err)
		}

		f.End = end.UTC()

		if f.Sta.Minute() != 0 {
			panic("--sta must not specify minutes, only full hours are supported")
		}

		if f.Sta.Second() != 0 {
			panic("--sta must not specify seconds, only full hours are supported")
		}
	}

	if !f.Sta.Before(f.End) {
		panic("--sta must be before --end")
	}
}
