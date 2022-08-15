package tra

import (
	"fmt"
	"time"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/framer"
)

func (r *run) newfra() framer.Frames {
	var err error

	var sta time.Time
	if !r.flags.Time.IsZero() {
		sta = r.flags.Time
	} else {
		sta = r.frasta()
	}

	if sta.After(time.Now().UTC()) {
		sta = r.frasta()
	}

	var end time.Time
	if r.flags.Duration != 0 {
		end = sta.Add(r.flags.Duration)
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

	var hfr []framer.Frame
	{
		hfr = fra.Exa().Dur(time.Hour)
	}

	return hfr
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

	var tra *trades.Trade
	{
		tra, err = r.storage.Latest()
		if err != nil {
			panic(err)
		}
	}

	if tra.Empty() {
		panic(fmt.Sprintf("trade %s must not be empty", tra.TS))
	}

	return tra.TS.AsTime()
}
