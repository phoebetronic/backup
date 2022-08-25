package apicliftx

import (
	"sort"
	"strings"
	"time"

	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/budget/v3"
	"github.com/xh3b4sd/budget/v3/pkg/breaker"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (f *FTX) Search(sta time.Time, end time.Time) []*trades.Trade {
	var err error

	var bre budget.Interface
	{
		c := breaker.Config{
			Failure: breaker.Failure{
				Budget: 10,
				Cooler: 10 * time.Second,
			},
			Timeout: breaker.Timeout{
				Action: 10 * time.Second,
				Budget: 10,
			},
		}

		bre, err = breaker.New(c)
		if err != nil {
			panic(err)
		}
	}

	var all []*trades.Trade
	for {
		var tra []*trades.Trade
		{
			act := func() error {
				tra, err = f.search(sta, end)
				if err != nil {
					return tracer.Mask(err)
				}

				return nil
			}

			err = bre.Execute(act)
			if err != nil {
				panic(err)
			}
		}

		{
			all = append(all, tra...)
		}

		if len(tra) < 5000 {
			break
		}

		{
			end = trafir(tra).TS.AsTime().Add(-1)
		}
	}

	sort.SliceStable(all, func(i, j int) bool {
		return all[i].TS.AsTime().UnixNano() < all[j].TS.AsTime().UnixNano()
	})

	return all
}

func (f *FTX) search(sta time.Time, end time.Time) ([]*trades.Trade, error) {
	var err error

	var res *markets.ResponseForTrades
	{
		req := &markets.RequestForTrades{
			ProductCode: f.market.Ass() + "-perp",
			Start:       sta.Unix(),
			End:         end.Unix(),
		}

		res, err = f.client.Trades(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var tra []*trades.Trade
	for _, r := range *res {
		t := &trades.Trade{}
		{
			t.LI = r.Liquidation
			t.PR = float32(r.Price)
			t.TS = timestamppb.New(r.Time)
		}

		if strings.ToLower(r.Side) == "buy" {
			t.LO = float32(r.Size)
		}

		if strings.ToLower(r.Side) == "sell" {
			t.SH = float32(r.Size)
		}

		{
			tra = append(tra, t)
		}
	}

	return tra, nil
}

func trafir(tra []*trades.Trade) *trades.Trade {
	fir := tra[0]

	for _, t := range tra {
		if t.TS.AsTime().Before(fir.TS.AsTime()) {
			fir = t
		}
	}

	return fir
}
