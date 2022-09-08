package apicliftx

import (
	"github.com/phoebetron/trades/typ/market"
)

func Default(ass string) *FTX {
	return New(Config{
		Mar: market.New(market.Config{
			Exc: "ftx",
			Ass: ass,
			Dur: 1,
		}),
	})
}
