package upl

import (
	"time"

	"github.com/spf13/cobra"
)

type flags struct {
	Exc string
	Ass string
	Kin string
	Tim time.Time
	tim string
	Fix bool
}

func (f *flags) Create(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Exc, "exc", "e", "", "The exchange to upload trades for, e.g. ftx.")
	cmd.Flags().StringVarP(&f.Ass, "ass", "a", "", "The asset to upload trades for, e.g. eth.")
	cmd.Flags().StringVarP(&f.Kin, "kin", "k", "", "The kind of backup data to validate, e.g. ord, tra.")
	cmd.Flags().StringVarP(&f.tim, "tim", "t", "", "Time string for backup data start date in form of yy-mm-dd.")
	cmd.Flags().BoolVarP(&f.Fix, "fix", "f", false, "Whether to fix broken trades like outlier prices.")
}

func (f *flags) Verify() {
	if f.Exc == "" {
		panic("-e/--exc must not be empty")
	}

	if f.Ass == "" {
		panic("-a/--ass must not be empty")
	}

	if f.Kin == "" {
		panic("-k/--kin must not be empty")
	}
	if f.Kin != "ord" && f.Kin != "tra" {
		panic("-k/--kin must be either ord or tra")
	}

	if f.tim == "" {
		panic("-t/--tim must not be empty")
	}

	if f.Kin == "ord" {
		tim, err := time.Parse("06-01-02T15:04:05", f.tim)
		if err != nil {
			panic(err)
		}

		f.Tim = tim.UTC()

		if f.Tim.Minute() != 0 || f.Tim.Second() != 0 {
			panic("-t/--tim must neither specify minutes nor seconds")
		}
	}

	if f.Kin == "tra" {
		tim, err := time.Parse("06-01-02", f.tim)
		if err != nil {
			panic(err)
		}

		f.Tim = tim.UTC()

		if f.Tim.Day() != 1 {
			panic("-t/--tim must specify the first day of the month")
		}
	}
}
