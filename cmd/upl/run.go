package upl

import (
	"bytes"
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
	"github.com/xh3b4sd/framer"
	"github.com/xh3b4sd/redigo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	var fra *framer.Framer
	if r.flags.Kin == "ord" {
		fra = framer.New(framer.Config{
			Sta: sta,
			End: end,
			Dur: time.Minute,
		})
	}
	if r.flags.Kin == "tra" {
		fra = framer.New(framer.Config{
			Sta: sta,
			End: end,
			Dur: time.Hour,
		})
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("starting backup between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	{
		defer fmt.Printf("finished backup\n")
	}

	var byt []byte
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

		var all []*orders.Bundle
		for _, h := range fra.List() {
			var ord *orders.Orders
			{
				ord, err = sto.Search(h.Sta)
				if err != nil {
					panic(err)
				}
			}

			{
				all = append(all, ord.BU...)
			}
		}

		ord := &orders.Orders{}
		{
			ord.EX = sto.Market().Exc()
			ord.AS = sto.Market().Ass()
			ord.ST = timestamppb.New(sta)
			ord.EN = timestamppb.New(end)
			ord.BU = all
		}

		{
			byt, err = proto.Marshal(ord)
			if err != nil {
				panic(err)
			}
		}

		defer func() {
			{
				fmt.Printf("removing orders between %s and %s\n", scrfmt(sta), scrfmt(end))
			}

			for _, h := range fra.List() {
				err := sto.Delete(h.Sta)
				if err != nil {
					panic(err)
				}
			}
		}()
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

		var all []*trades.Trade
		for _, h := range fra.List() {
			var tra *trades.Trades
			{
				tra, err = sto.Search(h.Sta)
				if err != nil {
					panic(err)
				}
			}

			{
				all = append(all, tra.TR...)
			}
		}

		tra := &trades.Trades{}
		{
			tra.EX = sto.Market().Exc()
			tra.AS = sto.Market().Ass()
			tra.ST = timestamppb.New(sta)
			tra.EN = timestamppb.New(end)
			tra.TR = all
		}

		{
			byt, err = proto.Marshal(tra)
			if err != nil {
				panic(err)
			}
		}

		defer func() {
			{
				fmt.Printf("removing trades between %s and %s\n", scrfmt(sta), scrfmt(end))
			}

			for _, h := range fra.List() {
				err := sto.Delete(h.Sta)
				if err != nil {
					panic(err)
				}
			}
		}()
	}

	// --------------------------------------------------------------------- //

	var rea bytes.Reader
	{
		rea = *bytes.NewReader(byt)
	}

	{
		fmt.Printf("buffered %s\n", r.siz(rea.Size()))
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

	{
		err := r.client.Upload(buc, filepath.Join(pre, suf), rea)
		if err != nil {
			panic(err)
		}
	}
}
