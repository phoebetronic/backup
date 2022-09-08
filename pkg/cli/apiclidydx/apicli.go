package apiclidydx

import (
	"github.com/phoebetron/dydxv3/client"
	"github.com/phoebetron/trades/typ/market"
)

type DyDx struct {
	cli *client.Client
	mar market.Market
}

func New(con Config) *DyDx {
	{
		con.Verify()
	}

	var cli *client.Client
	{
		cli = client.New(client.Config{})
	}

	return &DyDx{
		cli: cli,
		mar: con.Mar,
	}
}
