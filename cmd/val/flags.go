package val

import (
	"time"

	"github.com/spf13/cobra"
)

type flags struct {
	Exchange string
	Asset    string
	Kin      string
	Time     time.Time
	time     string
}

func (f *flags) Create(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Exchange, "exc", "e", "", "The exchange to validate trades for, e.g. ftx.")
	cmd.Flags().StringVarP(&f.Asset, "ass", "a", "", "The asset to validate trades for, e.g. eth.")
	cmd.Flags().StringVarP(&f.Kin, "kin", "k", "", "The kind of backup data to validate, e.g. ord, tra.")
	cmd.Flags().StringVarP(&f.time, "tim", "t", "", "Time string for backup data start date in form of yy-mm-dd.")
}

func (f *flags) Verify() {
	if f.Exchange == "" {
		panic("-e/--exc must not be empty")
	}

	if f.Asset == "" {
		panic("-a/--ass must not be empty")
	}

	if f.Kin == "" {
		panic("-k/--kin must not be empty")
	}
	if f.Kin != "ord" && f.Kin != "tra" {
		panic("-k/--kin must be either ord or tra")
	}

	if f.time == "" {
		panic("-t/--tim must not be empty")
	}

	if f.Kin == "ord" {
		tim, err := time.Parse("06-01-02T15:04:05", f.time)
		if err != nil {
			panic(err)
		}

		f.Time = tim.UTC()

		if f.Time.Minute() != 0 || f.Time.Second() != 0 {
			panic("-t/--tim must neither specify minutes nor seconds")
		}
	}

	if f.Kin == "tra" {
		tim, err := time.Parse("06-01-02", f.time)
		if err != nil {
			panic(err)
		}

		f.Time = tim.UTC()

		if f.Time.Day() != 1 {
			panic("-t/--tim must specify the first day of the month")
		}
	}
}
