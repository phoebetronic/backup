package apicliftx

import (
	"github.com/phoebetron/backup/pkg/fac/clifacftx"
	"github.com/phoebetron/backup/pkg/mis/env"
	"github.com/phoebetron/trades/typ/market"
)

func Default(ass string) *FTX {
	var err error

	var e env.Env
	{
		e = env.Create()
	}

	var f *clifacftx.FTX
	{
		c := clifacftx.Config{
			Key: e.FTX.ApiKey,
			Sec: e.FTX.ApiSecret,
		}

		f, err = clifacftx.New(c)
		if err != nil {
			panic(err)
		}
	}

	return New(Config{
		Cli: f.New(),
		Mar: market.New(market.Config{
			Exc: "ftx",
			Ass: ass,
			Dur: 1,
		}),
	})
}
