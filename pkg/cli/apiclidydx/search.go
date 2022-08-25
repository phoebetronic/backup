package apiclidydx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/budget/v3"
	"github.com/xh3b4sd/budget/v3/pkg/breaker"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	urlfmt = "https://api.dydx.exchange/v3/trades/%s-USD?startingBeforeOrAt=%s&limit=100"
)

func (d *DyDx) Search(sta time.Time, end time.Time) []*trades.Trade {
	var err error

	var bre budget.Interface
	{
		c := breaker.Config{
			Failure: breaker.Failure{
				Budget: 10,
				Cooler: 10 * time.Second,
			},
			Timeout: breaker.Timeout{
				Action: 10 * time.Second,
				Budget: 10,
			},
		}

		bre, err = breaker.New(c)
		if err != nil {
			panic(err)
		}
	}

	var all []*trades.Trade
	for {
		var tra []*trades.Trade
		{
			act := func() error {
				tra, err = d.search(sta, end)
				if err != nil {
					return tracer.Mask(err)
				}

				return nil
			}

			err = bre.Execute(act)
			if err != nil {
				panic(err)
			}
		}

		{
			all = append(all, tra...)
		}

		if len(tra) < 100 {
			break
		}

		{
			end = trafir(tra).TS.AsTime().Add(-1)
		}
	}

	sort.SliceStable(all, func(i, j int) bool {
		return all[i].TS.AsTime().UnixNano() < all[j].TS.AsTime().UnixNano()
	})

	return all
}

func (d *DyDx) search(sta time.Time, end time.Time) ([]*trades.Trade, error) {
	var err error

	var byt []byte
	{
		byt, err = d.byt(end)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res Response
	{
		err = json.Unmarshal(byt, &res)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var tra []*trades.Trade
	for _, v := range res.Trades {
		if v.CreatedAt.Before(sta) {
			continue
		}

		t := &trades.Trade{}
		{
			t.LI = v.Liquidation
			t.PR = musf32(v.Price)
			t.TS = timestamppb.New(v.CreatedAt)
		}

		if strings.ToLower(v.Side) == "buy" {
			t.LO = musf32(v.Size)
		}

		if strings.ToLower(v.Side) == "sell" {
			t.SH = musf32(v.Size)
		}

		{
			tra = append(tra, t)
		}
	}

	return tra, nil
}

func (d *DyDx) byt(end time.Time) ([]byte, error) {
	var err error

	var url string
	{
		url = fmt.Sprintf(urlfmt, strings.ToUpper(d.market.Ass()), end.Format(time.RFC3339))
	}

	var req *http.Request
	{
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *http.Response
	{
		res, err = d.client.Do(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}
		defer res.Body.Close()
	}

	var byt []byte
	{
		byt, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return byt, nil
}

func musf32(s string) float32 {
	f, e := strconv.ParseFloat(s, 32)
	if e != nil {
		panic(e)
	}

	return float32(f)
}

func trafir(tra []*trades.Trade) *trades.Trade {
	fir := tra[0]

	for _, t := range tra {
		if t.TS.AsTime().Before(fir.TS.AsTime()) {
			fir = t
		}
	}

	return fir
}
