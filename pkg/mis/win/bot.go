package win

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

//
//     tr are all trades available within which we are searching for windows
//     le is the desired time denominated window length we want to capture
//
func Bot(tr []*trades.Trade, le time.Duration) []Window {
	var s int
	var w []Window
	for {
		{
			w = append(w, Window{LE: &trades.Trades{}, RI: &trades.Trades{}})
		}

		var c int
		{
			c = len(w) - 1
		}

		var e int
		for i := range tr[s:] {
			var j int
			{
				j = s + i
			}

			if dur(tr[s], tr[j]) > le/2 {
				break
			}

			{
				w[c].LE.TR = append(w[c].LE.TR, tr[j])
			}

			{
				e = j
			}
		}

		for i := range tr[e:] {
			var j int
			{
				j = s + i
			}

			if dur(tr[e], tr[j]) > le/2 {
				break
			}

			{
				w[c].RI.TR = append(w[c].RI.TR, tr[j])
			}

			var de float32
			{
				de = del(tr[e], tr[j])
			}

			if de < w[c].BD {
				w[c].BI = i
				w[c].BD = de
			}

			if de > w[c].TD {
				w[c].TI = i
				w[c].TD = de
			}
		}

		{
			s = rig(tr, s, 5*time.Second)
		}

		if e+1 >= len(tr) {
			break
		}
	}

	return w
}
