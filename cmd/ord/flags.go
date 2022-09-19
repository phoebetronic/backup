package ord

import (
	"github.com/spf13/cobra"
)

type flags struct {
	Exchange string
	Asset    string
}

func (f *flags) Create(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Exchange, "exc", "e", "", "The exchange to backup trades from, e.g. ftx.")
	cmd.Flags().StringVarP(&f.Asset, "ass", "a", "", "The asset to backup trades for, e.g. eth.")
}

func (f *flags) Verify() {
	if f.Exchange == "" {
		panic("-e/--exc must not be empty")
	}

	if f.Asset == "" {
		panic("-a/--ass must not be empty")
	}
}
