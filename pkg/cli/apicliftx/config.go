package apicliftx

import (
	"github.com/go-numb/go-ftx/rest"
	"github.com/phoebetron/trades/typ/key"
)

type Config struct {
	Client *rest.Client
	Market *key.Key
}

func (c Config) Verify() {
	if c.Client == nil {
		panic("Market must not be empty")
	}
	if c.Market == nil {
		panic("Market must not be empty")
	}
}
