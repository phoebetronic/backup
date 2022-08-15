package val

import (
	"github.com/phoebetron/trades/typ/key"
)

func (r *run) newkey() *key.Key {
	var err error

	var k *key.Key
	{
		c := key.Config{
			Exc: r.flags.Exchange,
			Ass: "eth",
		}

		k, err = key.New(c)
		if err != nil {
			panic(err)
		}
	}

	return k
}
