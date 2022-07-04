package apicliftx

import (
	"sort"
	"time"

	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/budget/v3"
	"github.com/xh3b4sd/budget/v3/pkg/breaker"
	"github.com/xh3b4sd/tracer"
)

func (f *FTX) Search(sta time.Time, end time.Time) []trades.Trade {
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

	var all []trades.Trade
	for {
		var tra []trades.Trade
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
			end = trafir(tra).TS.Add(-1)
		}
	}

	sort.SliceStable(all, func(i, j int) bool {
		return all[i].TS.Unix() < all[j].TS.Unix()
	})

	return all

}

func (f *FTX) search(sta time.Time, end time.Time) ([]trades.Trade, error) {
	var tra []trades.Trade
	{
		req := &markets.RequestForTrades{
			ProductCode: f.market.Ass() + "-perp",
			Start:       sta.Unix(),
			End:         end.Unix(),
		}

		res, err := f.client.Trades(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		for _, r := range *res {
			var t trades.Trade
			{
				t.LI = r.Liquidation
				t.PR = r.Price
				t.TS = r.Time
			}

			if r.Side == "buy" {
				t.LO = r.Size
			}

			if r.Side == "sell" {
				t.SH = r.Size
			}

			{
				tra = append(tra, t)
			}
		}
	}

	return tra, nil
}

func trafir(tra []trades.Trade) trades.Trade {
	fir := tra[0]

	for _, t := range tra {
		if t.TS.Before(fir.TS) {
			fir = t
		}
	}

	return fir
}
