package apiclidydx

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/phoebetron/dydxv3/client/public/orderbook"
	"github.com/phoebetron/trades/typ/orders"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (d *DyDx) Orders() *orders.Bundle {
	var err error

	var req orderbook.GetRequest
	{
		req = orderbook.GetRequest{
			Market: fmt.Sprintf("%s-USD", strings.ToUpper(d.mar.Ass())),
		}
	}

	var res orderbook.GetResponse
	{
		res, err = d.cli.Pub.Ord.Get(req)
		if err != nil {
			tracer.Panic(err)
		}
	}

	{
		sort.SliceStable(res.Asks, func(i, j int) bool { return res.Asks[i].Pri() < res.Asks[j].Pri() })
		sort.SliceStable(res.Bids, func(i, j int) bool { return res.Bids[i].Pri() < res.Bids[j].Pri() })
	}

	var ask []*orders.Order
	for _, x := range res.Asks {
		ask = append(ask, &orders.Order{PR: x.Pri(), SI: x.Siz()})
	}

	var bid []*orders.Order
	for _, x := range res.Bids {
		bid = append(bid, &orders.Order{PR: x.Pri(), SI: x.Siz()})
	}

	var bun *orders.Bundle
	{
		bun = &orders.Bundle{
			AK: ask,
			BD: bid,
			TS: timestamppb.New(time.Now().UTC()),
		}
	}

	return bun
}
