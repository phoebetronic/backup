package ord

import (
	"github.com/phoebetron/trades/sto/ordersredis"
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/orders"
	"github.com/xh3b4sd/redigo"
)

func (r *run) newsto() orders.Storage {
	return ordersredis.New(ordersredis.Config{
		Mar: market.New(market.Config{
			Exc: r.flags.Exchange,
			Ass: r.flags.Asset,
			Dur: 1,
		}),
		Sor: redigo.Default().Sorted(),
	})
}
