package win

import "github.com/phoebetron/trades/typ/trades"

type Window struct {
	// BI is the bottom index of this window. This index points to the right
	// side trade with the lowest price.
	BI int
	// BD is the bottom delta of this window. This is the lower fraction of the
	// right side price change.
	BD float32
	// TI is the top index of this window. This index points to the right side
	// trade with the highest price.
	TI int
	// TD is the top delta of this window. This is the upper fraction of the
	// right side price change.
	TD float32
	LE *trades.Trades
	RI *trades.Trades
}
