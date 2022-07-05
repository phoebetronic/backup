package tra

import (
	"fmt"

	"github.com/phoebetron/backup/pkg/cli/apicliftx"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type run struct {
	cliftx *apicliftx.FTX
	cmdfla *fla
	misfra framer.Frames
	stotra trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	{
		r.cmdfla.Verify()
	}

	// --------------------------------------------------------------------- //

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
			fmt.Printf("fetching trades between %s and %s\n", scrfmt(h.Sta), scrfmt(h.End))
		}

		tra := &trades.Trades{}
		{
			tra.EX = r.stotra.Market().Exc()
			tra.AS = r.stotra.Market().Ass()
			tra.ST = timestamppb.New(h.Sta)
			tra.EN = timestamppb.New(h.End)
			tra.TR = r.cliftx.Search(h.Sta, h.End)
		}

		if len(tra.TR) == 0 {
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
