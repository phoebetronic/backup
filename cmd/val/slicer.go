package val

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

type Slicer struct {
	his time.Duration
	lis []*trades.Trade
}

func (s *Slicer) Add(t *trades.Trade) {
	{
		s.lis = append(s.lis, t)
	}

	for i := 0; i < len(s.lis); i++ {
		d := t.TS.AsTime().Sub(s.lis[i].TS.AsTime())
		if d > s.his {
			{
				copy(s.lis[i:], s.lis[i+1:])
				s.lis[len(s.lis)-1] = nil
				s.lis = s.lis[:len(s.lis)-1]
			}

			{
				i--
			}
		} else {
			break
		}
	}
}

func (s *Slicer) Lis() []*trades.Trade {
	return s.lis
}

func (s *Slicer) Avg() float32 {
	var sum float32

	for _, x := range s.lis {
		sum += x.PR
	}

	return sum / float32(len(s.lis))
}
