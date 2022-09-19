package val

import (
	"fmt"
	"time"

	"github.com/phoebetron/trades/sto/ordersredis"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/orders"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/redigo"
)

type run struct {
	flags *flags
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	{
		sta = r.flags.Time
	}

	var end time.Time
	{
		end = sta.AddDate(0, 1, 0)
	}

	// --------------------------------------------------------------------- //

	if r.flags.Kin == "ord" {
		var sto orders.Storage
		{
			sto = ordersredis.New(ordersredis.Config{
				Mar: market.New(market.Config{
					Exc: r.flags.Exchange,
					Ass: r.flags.Asset,
					Dur: 1,
				}),
				Sor: redigo.Default().Sorted(),
			})
		}

		{
			fmt.Printf("checking orders between %s and %s\n", scrfmt(sta), scrfmt(end))
		}

		var ord *orders.Orders
		{
			ord, err = sto.Search(sta)
			if err != nil {
				panic(err)
			}
		}

		{
			fmt.Printf("EX:    %s\n", ord.EX)
			fmt.Printf("AS:    %s\n", ord.AS)
			fmt.Printf("ST:    %s\n", scrfmt(ord.ST.AsTime()))
			fmt.Printf("EN:    %s\n", scrfmt(ord.EN.AsTime()))
			fmt.Printf("BU:    %d\n", len(ord.BU))
			fmt.Printf("FI:    %s\n", scrfmt(ord.BU[0].TS.AsTime()))
			fmt.Printf("LA:    %s\n", scrfmt(ord.BU[len(ord.BU)-1].TS.AsTime()))
		}
	}

	if r.flags.Kin == "tra" {
		var sto trades.Storage
		{
			sto = tradesredis.New(tradesredis.Config{
				Mar: market.New(market.Config{
					Exc: r.flags.Exchange,
					Ass: r.flags.Asset,
					Dur: 1,
				}),
				Sor: redigo.Default().Sorted(),
			})
		}

		{
			fmt.Printf("checking trades between %s and %s\n", scrfmt(sta), scrfmt(end))
		}

		var tra *trades.Trades
		{
			tra, err = sto.Search(sta)
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
}
