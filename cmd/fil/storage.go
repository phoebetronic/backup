package fil

import (
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/redigo"
)

func (r *run) newsto() trades.Storage {
	return tradesredis.New(tradesredis.Config{
		Mar: market.New(market.Config{
			Exc: r.flags.Exc,
			Ass: r.flags.Ass,
			Dur: 1,
		}),
		Sor: redigo.Default().Sorted(),
	})
}
