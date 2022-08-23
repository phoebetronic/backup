package apiclidydx

import (
	"github.com/phoebetron/trades/typ/key"
)

func Default(ass string) *DyDx {
	var err error

	var m *key.Key
	{
		c := key.Config{
			Exc: "dydx",
			Ass: ass,
		}

		m, err = key.New(c)
		if err != nil {
			panic(err)
		}
	}

	var dydx *DyDx
	{
		c := Config{
			Market: m,
		}

		dydx = New(c)
	}

	return dydx
}
