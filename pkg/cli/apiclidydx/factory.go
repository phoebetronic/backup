package apiclidydx

import "github.com/phoebetron/trades/typ/market"

func Default(ass string) *DyDx {
	return New(Config{
		Mar: market.New(market.Config{
			Exc: "dydx",
			Ass: ass,
			Dur: 1,
		}),
	})
}
