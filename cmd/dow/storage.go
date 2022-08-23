package dow

import (
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/key"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/redigo"
)

func (r *run) newsto() trades.Storage {
	var err error

	var k *key.Key
	{
		c := key.Config{
			Exc: r.flags.Exchange,
			Ass: r.flags.Asset,
		}

		k, err = key.New(c)
		if err != nil {
			panic(err)
		}
	}

	var red trades.Storage
	{
		c := tradesredis.Config{
			Key: k,
			Sor: redigo.Default().Sorted(),
		}

		red, err = tradesredis.New(c)
		if err != nil {
			panic(err)
		}
	}

	return red
}
