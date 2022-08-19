package upl

import (
	"time"

	"github.com/xh3b4sd/framer"
)

func (r *run) newfra() framer.Frames {
	var sta time.Time
	{
		sta = r.flags.Time
	}

	var end time.Time
	{
		end = sta.AddDate(0, 1, 0)
	}

	var fra *framer.Framer
	{
		fra = framer.New(framer.Config{
			Sta: sta,
			End: end,
			Dur: time.Hour,
		})
	}

	return fra.List()
}
