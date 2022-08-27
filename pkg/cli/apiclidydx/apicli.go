package apiclidydx

import (
	"net/http"

	"github.com/phoebetron/trades/typ/market"
)

type DyDx struct {
	cli *http.Client
	mar *market.Market
}

func New(con Config) *DyDx {
	{
		con.Verify()
	}

	return &DyDx{
		cli: &http.Client{},
		mar: con.Mar,
	}
}
