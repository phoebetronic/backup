package ens

import (
	"time"

	"github.com/phoebetron/series/buck"
	"github.com/phoebetron/series/buff"
	"github.com/phoebetron/trades/typ/trades"
)

func (r *run) num(fra []*trades.Trades, con buff.Conf) int {
	var tra []*trades.Trade
	for _, f := range fra {
		tra = append(tra, f.TR...)
	}

	var m map[string][]buck.Wndw
	{
		m = buck.Buck(tra, buck.Prec(), time.Hour, con)
	}

	var num int
	{
		num = len(m[con.Hash()][0].SE)
	}

	return num
}
