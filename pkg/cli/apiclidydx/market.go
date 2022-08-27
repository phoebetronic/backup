package apiclidydx

import "github.com/phoebetron/trades/typ/market"

func (d DyDx) Market() *market.Market {
	return d.mar
}
