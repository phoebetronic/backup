package ord

import (
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/orders"
)

type Client interface {
	Market() market.Market
	Orders() *orders.Bundle
}
