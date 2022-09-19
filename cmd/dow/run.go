package dow

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/trades/sto/ordersredis"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/orders"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/redigo"
	"google.golang.org/protobuf/proto"
)

type run struct {
	client *apicliaws.AWS
	flags  *flags
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.client = apicliaws.New()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	var end time.Time
	if r.flags.Kin == "ord" {
		sta = r.flags.Time
		end = sta.Add(time.Hour)
	}
	if r.flags.Kin == "tra" {
		sta = r.flags.Time
		end = sta.AddDate(0, 1, 0)
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("checking backup between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	var buc string
	{
		buc = "xh3b4sd-phoebe-backup"
	}

	var pre string
	if r.flags.Kin == "ord" {
		pre = fmt.Sprintf("ord-raw.exc-%s.ass-%s", r.flags.Exchange, r.flags.Asset)
	}
	if r.flags.Kin == "tra" {
		pre = fmt.Sprintf("tra-raw.exc-%s.ass-%s", r.flags.Exchange, r.flags.Asset)
	}

	var suf string
	if r.flags.Kin == "ord" {
		suf = fmt.Sprintf("%s.pb.raw", ordfmt(sta))
	}
	if r.flags.Kin == "tra" {
		suf = fmt.Sprintf("%s.pb.raw", trafmt(sta))
	}

	var byt []byte
	{
		byt, err = r.client.Download(buc, filepath.Join(pre, suf))
		if err != nil {
			panic(err)
		}
	}

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

		ord := &orders.Orders{}
		{
			err := proto.Unmarshal(byt, ord)
			if err != nil {
				panic(err)
			}
		}

		{
			err := sto.Create(ord.ST.AsTime(), ord)
			if tradesredis.IsAlreadyExists(err) {
				err = sto.Update(ord.ST.AsTime(), ord)
				if err != nil {
					panic(err)
				}
			} else if err != nil {
				panic(err)
			}
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

		tra := &trades.Trades{}
		{
			err := proto.Unmarshal(byt, tra)
			if err != nil {
				panic(err)
			}
		}

		{
			err := sto.Create(tra.ST.AsTime(), tra)
			if tradesredis.IsAlreadyExists(err) {
				err = sto.Update(tra.ST.AsTime(), tra)
				if err != nil {
					panic(err)
				}
			} else if err != nil {
				panic(err)
			}
		}
	}
}
