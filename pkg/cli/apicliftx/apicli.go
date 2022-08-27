package apicliftx

import (
	"github.com/go-numb/go-ftx/rest"
	"github.com/phoebetron/trades/typ/market"
)

type FTX struct {
	cli *rest.Client
	mar *market.Market
}

func New(con Config) *FTX {
	{
		con.Verify()
	}

	return &FTX{
		cli: con.Cli,
		mar: con.Mar,
	}
}
