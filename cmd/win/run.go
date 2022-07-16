package win

import (
	"fmt"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/backup/pkg/mis/win"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
)

type run struct {
	cliaws *apicliaws.AWS
	cmdfla *fla
	misfra framer.Frames
	stotra trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.cmdfla.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.cliaws = apicliaws.Default()
	}

	{
		r.stotra = tradesredis.Default()
	}

	{
		r.misfra = r.franew()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	var end time.Time
	{
		sta = r.misfra.Min().Sta
		end = r.misfra.Max().End
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("creating frames between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	now := time.Now()

	var tra *trades.Trades
	{
		tra, err = r.stotra.Search(sta)
		if err != nil {
			panic(err)
		}
	}

	var w []win.Window
	{
		w = win.Bot(tra.TR, 60*time.Minute)
	}

	if r.cmdfla.CSV {
		r.csv(w)
	}

	if r.cmdfla.Dra {
		r.dra(w)
	}

	{
		fmt.Printf("produced %d window frames within %s\n", len(w), time.Since(now).Round(10*time.Millisecond))
	}
}
