package win

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

func dur(a *trades.Trade, b *trades.Trade) time.Duration {
	return b.TS.AsTime().Sub(a.TS.AsTime())
}

func del(a *trades.Trade, b *trades.Trade) float32 {
	return (b.PR - a.PR) / a.PR
}

func rig(tr []*trades.Trade, i int, le time.Duration) int {
	for p := i; p < len(tr); p++ {
		if dur(tr[i], tr[p]) > le {
			return p
		}
	}

	return len(tr) - 1
}
