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

	var m map[string][]buck.Wndw
	{
		m = buck.Buck(tra, buck.Prec(), time.Hour, r.miscon)
	}

	var num int
	{
		num = len(m[r.miscon.Hash()][0].SE)
	}

	return num
}
