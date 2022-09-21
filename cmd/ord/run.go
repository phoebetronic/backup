package ord

import (
	"fmt"
	"time"

	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/orders"
	"github.com/phoebetron/trades/typ/orders/buffer"
	"github.com/spf13/cobra"
)

type run struct {
	client  Client
	flags   *flags
	storage orders.Storage
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

	var mar market.Market
	{
		mar = market.New(market.Config{
			Exc: r.flags.Exchange,
			Ass: r.flags.Asset,
			Dur: time.Minute,
		})
	}

	var buf buffer.Buffer
	{
		buf = buffer.New(buffer.Config{
			Mar: mar,
		})
	}

	go func() {
		for {
			var bun *orders.Bundle
			{
				bun = r.client.Orders()
			}

			{
				buf.Buffer(bun)
			}

			{
				time.Sleep(3 * time.Second)
			}
		}
	}()

	go func() {
		for {
			var cur time.Time
			{
				cur = time.Now().UTC()
			}

			var dur time.Duration
			{
				dur = cur.Truncate(mar.Dur()).Add(mar.Dur()).Sub(cur)
			}

			{
				time.Sleep(dur)
			}

			{
				buf.Finish(time.Now().UTC())
			}
		}
	}()

	for o := range buf.Orders() {
		{
			fmt.Printf(
				"creating %s orders from %s between %s and %s\n",
				r.client.Market().Ass(),
				r.client.Market().Exc(),
				scrfmt(o.ST.AsTime()),
				scrfmt(o.EN.AsTime()),
			)
		}

		for i := range o.BU {
			o.BU[i].MI = o.BU[i].Mid()
		}

		{
			err := r.storage.Create(o.ST.AsTime(), o)
			if tradesredis.IsAlreadyExists(err) {
				err = r.storage.Update(o.ST.AsTime(), o)
				if err != nil {
					panic(err)
				}
			} else if err != nil {
				panic(err)
			}
		}
	}
}
