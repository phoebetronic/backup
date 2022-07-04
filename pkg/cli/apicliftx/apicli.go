package apicliftx

import (
	"github.com/go-numb/go-ftx/rest"
	"github.com/phoebetron/trades/typ/key"
)

type FTX struct {
	client *rest.Client
	market *key.Key
}

func New(con Config) *FTX {
	{
		con.Verify()
	}

	f := &FTX{
		client: con.Client,
		market: con.Market,
	}

	return f
}
