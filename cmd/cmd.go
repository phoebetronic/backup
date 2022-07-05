package cmd

import (
	"github.com/phoebetron/backup/cmd/com"
	"github.com/phoebetron/backup/cmd/tra"
	"github.com/phoebetron/backup/cmd/upl"
	"github.com/phoebetron/backup/cmd/ver"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

var (
	use = "backup"
	sho = "Manage data backups."
	lon = "Manage data backups."
)

func New() (*cobra.Command, error) {
	var err error

	// --------------------------------------------------------------------- //

	var cmdCom *cobra.Command
	{
		c := com.Config{}

		cmdCom, err = com.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdTra *cobra.Command
	{
		c := tra.Config{}

		cmdTra, err = tra.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdUpl *cobra.Command
	{
		c := upl.Config{}

		cmdUpl, err = upl.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdVer *cobra.Command
	{
		c := ver.Config{}

		cmdVer, err = ver.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// --------------------------------------------------------------------- //

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
			// We slience errors because we do not want to see spf13/cobra printing.
			// The errors returned by the commands will be propagated to the main.go
			// anyway, where we have custom error printing for the command line
			// tool.
			SilenceErrors: true,
			SilenceUsage:  true,
		}
	}

	{
		c.SetHelpCommand(&cobra.Command{Hidden: true})
	}

	{
		c.AddCommand(cmdCom)
		c.AddCommand(cmdTra)
		c.AddCommand(cmdUpl)
		c.AddCommand(cmdVer)
	}

	return c, nil
}
