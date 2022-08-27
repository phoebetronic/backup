package apicliftx

import (
	"github.com/go-numb/go-ftx/rest"
	"github.com/phoebetron/trades/typ/market"
)

type Config struct {
	Cli *rest.Client
	Mar *market.Market
}

func (c Config) Verify() {
	if c.Cli == nil {
		panic("Config.Cli must not be empty")
	}
	if c.Mar == nil {
		panic("Config.Mar must not be empty")
	}
}
