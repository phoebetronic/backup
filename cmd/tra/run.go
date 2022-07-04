package tra

import (
	"fmt"

	"github.com/phoebetron/backup/pkg/cli/apicliftx"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
)

type run struct {
	cmdfla *fla
	cliftx *apicliftx.FTX
	misfra framer.Frames
	stotra trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	{
		r.cmdfla.Verify()
	}

	{
		r.cliftx = apicliftx.Default()
	}

	{
		r.stotra = tradesredis.Default()
	}

	{
		r.misfra = r.franew()
	}

	// --------------------------------------------------------------------- //

	for _, h := range r.misfra {
		{
			fmt.Printf("fetching trades between %s and %s\n", timfmt(h.Sta), timfmt(h.End))
		}

		var tra []trades.Trade
		{
			tra = r.cliftx.Search(h.Sta, h.End)
		}

		if len(tra) == 0 {
			continue
		}

		{
			err := r.stotra.Create(h.Sta, tra)
			if tradesredis.IsAlreadyExists(err) {
				err = r.stotra.Update(h.Sta, tra)
				if err != nil {
					panic(err)
				}
			} else if err != nil {
				panic(err)
			}
		}
	}
}
