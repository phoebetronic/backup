package fil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/trades/sto/ordersredis"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/orders"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
	"github.com/xh3b4sd/redigo"
	"google.golang.org/protobuf/proto"
)

type run struct {
	client  *apicliaws.AWS
	flags   *flags
	storage trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.client = apicliaws.New()
		r.storage = r.newsto()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	var end time.Time
	if r.flags.Kin == "ord" {
		sta = time.Date(
			r.flags.Sta.Year(),
			r.flags.Sta.Month(),
			r.flags.Sta.Day(),
			r.flags.Sta.Hour(),
			0,
			0,
			0,
			time.UTC,
		)
		end = sta.Add(time.Hour)
	}
	if r.flags.Kin == "tra" {
		sta = time.Date(
			r.flags.Sta.Year(),
			r.flags.Sta.Month(),
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		)
		end = sta.AddDate(0, 1, 0)
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("checking backup between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	var byt []byte
	if r.flags.Kin == "ord" {
		var sto orders.Storage
		{
			sto = ordersredis.New(ordersredis.Config{
				Mar: market.New(market.Config{
					Exc: r.flags.Exc,
					Ass: r.flags.Ass,
					Dur: 1,
				}),
				Sor: redigo.Default().Sorted(),
			})
		}

		var ord *orders.Orders
		{
			ord, err = sto.Search(sta)
			if ordersredis.IsNotFound(err) {
				var buc string
				{
					buc = "xh3b4sd-phoebe-backup"
				}

				var pre string
				{
					pre = fmt.Sprintf("ord-raw.exc-%s.ass-%s", r.flags.Exc, r.flags.Ass)
				}

				var suf string
				{
					suf = fmt.Sprintf("%s.pb.raw", ordfmt(sta))
				}

				var byt []byte
				{
					byt, err = r.client.Download(buc, filepath.Join(pre, suf))
					if err != nil {
						panic(err)
					}
				}

				ord = &orders.Orders{}
				{
					err := proto.Unmarshal(byt, ord)
					if err != nil {
						panic(err)
					}
				}

				{
					err := sto.Create(ord.ST.AsTime(), ord)
					if ordersredis.IsAlreadyExists(err) {
						err = sto.Update(ord.ST.AsTime(), ord)
						if err != nil {
							panic(err)
						}
					} else if err != nil {
						panic(err)
					}
				}
			} else if err != nil {
				panic(err)
			}
		}

		{
			fmt.Printf("creating frames with %s resolution\n", r.flags.Dur)
		}

		var fra *orders.Framer
		{
			fra = ord.Frame(framer.Config{
				Sta: r.flags.Sta,
				End: r.flags.End,
				Len: r.flags.Dur,
			})
		}

		var lis []*orders.Bundle
		for !fra.Last() {
			lis = append(lis, fra.Next().BU...)
		}

		{
			byt, err = json.Marshal(lis)
			if err != nil {
				panic(err)
			}
		}
	}

	if r.flags.Kin == "tra" {
		var sto trades.Storage
		{
			sto = tradesredis.New(tradesredis.Config{
				Mar: market.New(market.Config{
					Exc: r.flags.Exc,
					Ass: r.flags.Ass,
					Dur: 1,
				}),
				Sor: redigo.Default().Sorted(),
			})
		}

		var tra *trades.Trades
		{
			tra, err = sto.Search(sta)
			if tradesredis.IsNotFound(err) {
				var buc string
				{
					buc = "xh3b4sd-phoebe-backup"
				}

				var pre string
				{
					pre = fmt.Sprintf("tra-raw.exc-%s.ass-%s", r.flags.Exc, r.flags.Ass)
				}

				var suf string
				{
					suf = fmt.Sprintf("%s.pb.raw", trafmt(sta))
				}

				var byt []byte
				{
					byt, err = r.client.Download(buc, filepath.Join(pre, suf))
					if err != nil {
						panic(err)
					}
				}

				tra = &trades.Trades{}
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
			} else if err != nil {
				panic(err)
			}
		}

		{
			fmt.Printf("creating frames with %s resolution\n", r.flags.Dur)
		}

		var fra *trades.Framer
		{
			fra = tra.Frame(framer.Config{
				Sta: r.flags.Sta,
				End: r.flags.End,
				Len: r.flags.Dur,
			})
		}

		var lis []*trades.Trades
		for !fra.Last() {
			lis = append(lis, fra.Next())
		}

		{
			byt, err = json.Marshal(lis)
			if err != nil {
				panic(err)
			}
		}
	}

	{
		err := ioutil.WriteFile(r.flags.Pat, byt, 0644)
		if err != nil {
			panic(err)
		}
	}
}
