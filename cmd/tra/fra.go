package tra

import (
	"fmt"
	"time"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/framer"
)

func (r *run) franew() framer.Frames {
	var err error

	var sta time.Time
	if !r.cmdfla.Time.IsZero() {
		sta = r.cmdfla.Time
	} else {
		sta = r.frasta()
	}

	if sta.After(time.Now().UTC()) {
		sta = r.frasta()
	}

	var end time.Time
	if r.cmdfla.Duration != 0 {
		end = sta.Add(r.cmdfla.Duration)
	} else {
		end = r.fraend()
	}

	if end.After(time.Now().UTC()) {
		end = r.fraend()
	}

	var fra framer.Interface
	{
		c := framer.Config{
			Sta: sta,
			End: end,
		}

		fra, err = framer.New(c)
		if err != nil {
			panic(err)
		}
	}

	var dfr []framer.Frame
	{
		dfr = fra.Exa().Hour()
	}

	return dfr
}

func (r *run) fraend() time.Time {
	var end time.Time
	{
		end = time.Now().UTC().Truncate(time.Minute).Add(-1 * time.Minute)
	}

	return end
}

func (r *run) frasta() time.Time {
	var err error

	var tra trades.Trade
	{
		tra, err = r.stotra.Latest()
		if err != nil {
			panic(err)
		}
	}

	if tra.Empty() {
		panic(fmt.Sprintf("trade %s must not be empty", tra.TS))
	}

	return tra.TS
}
