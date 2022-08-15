package tra

import (
	"time"

	"github.com/phoebetron/trades/typ/key"
	"github.com/phoebetron/trades/typ/trades"
)

type Client interface {
	Market() *key.Key
	Search(sta time.Time, end time.Time) []*trades.Trade
}
