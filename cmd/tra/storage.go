package tra

import (
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/redigo"
)

func (r *run) newsto() trades.Storage {
	return tradesredis.New(tradesredis.Config{
		Mar: market.New(market.Config{
			Exc: r.flags.Exchange,
			Ass: r.flags.Asset,
			Dur: 1,
		}),
		Sor: redigo.Default().Sorted(),
	})
}
