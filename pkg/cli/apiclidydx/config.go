package apiclidydx

import (
	"github.com/phoebetron/trades/typ/key"
)

type Config struct {
	Market *key.Key
}

func (c Config) Verify() {
	if c.Market == nil {
		panic("Market must not be empty")
	}
}
