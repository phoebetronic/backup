package val

import (
	"time"

	"github.com/xh3b4sd/framer"
)

func (r *run) newfra() framer.Frames {
	var err error

	var sta time.Time
	{
		sta = r.flags.Time
	}

	var end time.Time
	{
		end = sta.AddDate(0, 1, 0)
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
