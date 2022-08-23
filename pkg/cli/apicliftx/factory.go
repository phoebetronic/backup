package apicliftx

import (
	"github.com/phoebetron/backup/pkg/fac/clifacftx"
	"github.com/phoebetron/backup/pkg/mis/env"
	"github.com/phoebetron/trades/typ/key"
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

	var m *key.Key
	{
		c := key.Config{
			Exc: "ftx",
			Ass: ass,
		}

		m, err = key.New(c)
		if err != nil {
			panic(err)
		}
	}

	var ftx *FTX
	{
		c := Config{
			Client: f.New(),
			Market: m,
		}

		ftx = New(c)
	}

	return ftx
}
