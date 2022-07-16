package win

import "github.com/phoebetron/trades/typ/trades"

type Window struct {
	IN int
	SI float32
	LE *trades.Trades
	RI *trades.Trades
}
