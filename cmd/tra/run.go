package tra

import (
	"fmt"

	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type run struct {
	client  Client
	flags   *flags
	storage trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.client = r.newcli()
		r.storage = r.newsto()
	}

	// --------------------------------------------------------------------- //

	var fra *framer.Framer
	{
		fra = r.newfra()
	}

	for !fra.Last() {
		var nex framer.Frame
		{
			nex = fra.Next()
		}

		{
			fmt.Printf(
				"fetching %s trades from %s between %s and %s\n",
				r.client.Market().Ass(),
				r.client.Market().Exc(),
				scrfmt(nex.Sta),
				scrfmt(nex.End),
			)
		}

		tra := &trades.Trades{}
		{
			tra.EX = r.storage.Market().Exc()
			tra.AS = r.storage.Market().Ass()
			tra.ST = timestamppb.New(nex.Sta)
			tra.EN = timestamppb.New(nex.End)
			tra.TR = r.client.Trades(nex.Sta, nex.End)
		}

		if len(tra.TR) == 0 {
			continue
		}

		{
			err := r.storage.Create(nex.Sta, tra)
			if tradesredis.IsAlreadyExists(err) {
				err = r.storage.Update(nex.Sta, tra)
				if err != nil {
					panic(err)
				}
			} else if err != nil {
				panic(err)
			}
		}
	}
}
