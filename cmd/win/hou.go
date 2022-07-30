package win

import (
	"fmt"
	"math"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/ind"
	"github.com/phoebetron/series/buck"
	"github.com/phoebetron/trades/typ/trades"
)

func (r *run) hou(tra *trades.Trades) {
	var n time.Time
	{
		n = time.Now()
	}

	var p string
	{
		p = filepath.Join("dat", r.miscon.Hash())
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
			s = append(s, v.SE)
		}

		var f string
		{
			f = fmt.Sprintf("%s.%.3f.csv", sgn(b), math.Abs(float64(b)))
		}

		var d string
		{
			d = filepath.Join(p, "csv", datfmt(tra.ST.AsTime()), clofmt(tra.ST.AsTime()))
		}

		var i ind.Item
		{
			i = ind.Item{
				Buc: b,
				Cou: len(s),
				Fil: filepath.Join(d, f),
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
		r.csv(s, i.Fil)
	}

	{
		r.dir(filepath.Join(p, "ind"))
		r.ind(filepath.Join(p, "ind", "ind.json"), l)
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
