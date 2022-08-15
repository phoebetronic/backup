package apicliftx

import (
	"github.com/phoebetron/trades/typ/key"
)

func (f FTX) Market() *key.Key {
	return f.market
}
