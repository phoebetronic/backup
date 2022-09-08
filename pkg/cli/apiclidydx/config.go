package apiclidydx

import (
	"github.com/phoebetron/trades/typ/market"
)

type Config struct {
	Mar market.Market
}

func (c Config) Verify() {
	if c.Mar == nil {
		panic("Config.Mar must not be empty")
	}
}
