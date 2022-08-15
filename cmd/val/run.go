package val

import (
	"fmt"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/trades/typ/key"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
)

type run struct {
	client  *apicliaws.AWS
	flags   *flags
	frames  framer.Frames
	key     *key.Key
	storage trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.key = r.newkey()
	}

	{
		r.client = apicliaws.Default()
	}

	{
		r.storage = r.newsto()
	}

	{
		r.frames = r.newfra()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	var end time.Time
	{
		sta = r.frames.Min().Sta
		end = r.frames.Max().End
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("checking backup between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	var tra *trades.Trades
	{
		tra, err = r.storage.Search(sta)
		if err != nil {
			panic(err)
		}
	}

	{
		fmt.Printf("EX:    %s\n", tra.EX)
		fmt.Printf("AS:    %s\n", tra.AS)
		fmt.Printf("ST:    %s\n", scrfmt(tra.ST.AsTime()))
		fmt.Printf("EN:    %s\n", scrfmt(tra.EN.AsTime()))
		fmt.Printf("TR:    %d\n", len(tra.TR))
		fmt.Printf("FI:    %s\n", scrfmt(tra.TR[0].TS.AsTime()))
		fmt.Printf("LA:    %s\n", scrfmt(tra.TR[len(tra.TR)-1].TS.AsTime()))
	}
}
