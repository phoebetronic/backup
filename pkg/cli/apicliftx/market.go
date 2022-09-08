package apicliftx

import "github.com/phoebetron/trades/typ/market"

func (f FTX) Market() market.Market {
	return f.mar
}
