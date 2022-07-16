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
			w = append(w, Window{})
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
				w[c].LE = append(w[c].LE, tr[j])
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
				w[c].RI = append(w[c].RI, tr[j])
			}

			var si float32
			{
				si = inc(tr[e], tr[j])
			}

			if si > w[c].SI {
				w[c].IN = i
				w[c].SI = si
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

func dur(a *trades.Trade, b *trades.Trade) time.Duration {
	return b.TS.AsTime().Sub(a.TS.AsTime())
}

func inc(a *trades.Trade, b *trades.Trade) float32 {
	return (b.PR - a.PR) / a.PR
}

func rig(tr []*trades.Trade, i int, le time.Duration) int {
	for p := i; p < len(tr); p++ {
		if dur(tr[i], tr[p]) > le {
			return p
		}
	}

	return len(tr) - 1
}
