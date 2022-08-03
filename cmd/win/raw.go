package win

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/ind"
	"github.com/phoebetron/series/buck"
	"github.com/phoebetron/trades/typ/trades"
)

func (r *run) raw(pat string, num int, tra *trades.Trades) {
	var n time.Time
	{
		n = time.Now()
	}

	var m map[float32][]buck.Wndw
	{
		m = buck.Buck(tra.TR, buck.Prec(), r.miscon, time.Hour)
	}

	var c int
	var l ind.List
	for b, w := range m {
		var s [][]string

		for _, v := range w {
			if len(v.SE) == num {
				s = append(s, v.SE)
			} else {
				fmt.Printf(
					"ignoring window of length %d due to strict requirement of window length %d\n",
					len(v.SE),
					num,
				)
			}
		}

		var n string
		{
			n = nam(b)
		}

		var d string
		{
			d = filepath.Join(pat, "raw", datfmt(tra.ST.AsTime()), clofmt(tra.ST.AsTime()))
		}

		var i ind.Item
		{
			i = ind.Item{
				Buc: b,
				Cou: len(s),
				Fil: filepath.Join(d, n),
				Tim: tra.ST.AsTime(),
			}
		}

		{
			c += i.Cou
		}

		{
			l = append(l, i)
		}

		r.dir(d)
		r.wri(s, i.Fil)
	}

	{
		r.dir(filepath.Join(pat, "ind"))
		r.ind(pat, l)
	}

	{
		fmt.Printf(
			"created %d window frames from %d trades between %s and %s within %s\n",
			c,
			len(tra.TR),
			scrfmt(tra.ST.AsTime()),
			scrfmt(tra.EN.AsTime()),
			time.Since(n).Round(10*time.Millisecond),
		)
	}
}
