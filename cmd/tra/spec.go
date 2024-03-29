package tra

import (
	"time"

	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/trades"
)

type Client interface {
	Market() market.Market
	Trades(sta time.Time, end time.Time) []*trades.Trade
}
