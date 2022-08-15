package tra

import (
	"fmt"

	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/key"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type run struct {
	client  Client
	flags   *flags
	frames  framer.Frames
	key     *key.Key
	storage trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.key = r.newkey()
	}

	{
		r.client = r.newcli()
	}

	{
		r.storage = r.newsto()
	}

	{
		r.frames = r.newfra()
	}

	// --------------------------------------------------------------------- //

	for _, h := range r.frames {
		{
			fmt.Printf(
				"fetching %s trades from %s between %s and %s\n",
				r.client.Market().Ass(),
				r.client.Market().Exc(),
				scrfmt(h.Sta),
				scrfmt(h.End),
			)
		}

		tra := &trades.Trades{}
		{
			tra.EX = r.storage.Market().Exc()
			tra.AS = r.storage.Market().Ass()
			tra.ST = timestamppb.New(h.Sta)
			tra.EN = timestamppb.New(h.End)
			tra.TR = r.client.Search(h.Sta, h.End)
		}

		if len(tra.TR) == 0 {
			continue
		}

		{
			err := r.storage.Create(h.Sta, tra)
			if tradesredis.IsAlreadyExists(err) {
				err = r.storage.Update(h.Sta, tra)
				if err != nil {
					panic(err)
				}
			} else if err != nil {
				panic(err)
			}
		}
	}
}
