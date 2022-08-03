package mod

import (
	"time"

	"github.com/phoebetron/series/buck"
	"github.com/phoebetron/trades/typ/trades"
)

func (r *run) num(fra []*trades.Trades) int {
	var tra []*trades.Trade
	for _, f := range fra {
		tra = append(tra, f.TR...)
	}

	var m map[float32][]buck.Wndw
	{
		m = buck.Buck(tra, buck.Prec(), r.miscon, time.Hour)
	}

	var num int
	for _, w := range m {
		num = len(w[0].SE)
	}

	return num
}
