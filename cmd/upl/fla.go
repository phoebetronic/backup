package upl

import (
	"github.com/spf13/cobra"
)

type fla struct {
	Con bool
	Dat string
	Fil string
}

func (f *fla) Create(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&f.Con, "con", "c", false, "Run a continuous process for uploading periodically.")
	cmd.Flags().StringVarP(&f.Dat, "dat", "d", "", "The data folder on the file system.")
	cmd.Flags().StringVarP(&f.Fil, "fil", "f", "dump.rdb", "The snapshot file on the file system.")
}

func (f *fla) Verify() {
	{
		if f.Dat == "" {
			panic("-d/--dat must not be empty")
		}
		if f.Fil == "" {
			panic("-f/--fil must not be empty")
		}
	}
}
