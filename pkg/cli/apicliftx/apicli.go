package apicliftx

import (
	"github.com/phoebetron/ftxapi/client"
	"github.com/phoebetron/trades/typ/market"
)

type FTX struct {
	cli *client.Client
	mar market.Market
}

func New(con Config) *FTX {
	{
		con.Verify()
	}

	var cli *client.Client
	{
		cli = client.New(client.Config{})
	}

	return &FTX{
		cli: cli,
		mar: con.Mar,
	}
}
