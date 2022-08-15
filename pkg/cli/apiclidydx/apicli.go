package apiclidydx

import (
	"net/http"

	"github.com/phoebetron/trades/typ/key"
)

type DyDx struct {
	client *http.Client
	market *key.Key
}

func New(con Config) *DyDx {
	{
		con.Verify()
	}

	f := &DyDx{
		client: &http.Client{},
		market: con.Market,
	}

	return f
}
