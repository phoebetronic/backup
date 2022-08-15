package dow

import (
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/redigo"
)

func (r *run) newsto() trades.Storage {
	var err error

	var red trades.Storage
	{
		c := tradesredis.Config{
			Key: r.key,
			Sor: redigo.Default().Sorted(),
		}

		red, err = tradesredis.New(c)
		if err != nil {
			panic(err)
		}
	}

	return red
}
