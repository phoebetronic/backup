package win

import (
	"fmt"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/series/buff"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
)

const (
	fratra = 0.70
	frates = 0.15
	fraval = 0.15
)

type run struct {
	cliaws *apicliaws.AWS
	cmdfla *fla
	miscon buff.Conf
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
		r.miscon = r.connew()
	}

	{
		r.misfra = r.franew()
	}

	{
		r.stotra = tradesredis.Default()
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
		fmt.Printf("fetching trades between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	var tra *trades.Trades
	{
		tra, err = r.stotra.Search(sta)
		if err != nil {
			panic(err)
		}
	}

	var fra []*trades.Trades
	{
		fra = tra.Frame(r.misfra)
	}

	for _, f := range fra {
		r.hou(f)
	}
}
