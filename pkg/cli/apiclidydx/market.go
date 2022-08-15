package apiclidydx

import (
	"github.com/phoebetron/trades/typ/key"
)

func (d DyDx) Market() *key.Key {
	return d.market
}
