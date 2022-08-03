package ind

import (
	"fmt"
	"strconv"
)

type Index struct {
	Lis List   `json:"lis"`
	Sta []Stat `json:"sta"`
}

func (i Index) EncI() map[string]string {
	m := map[string]string{}

	for i, s := range i.Sta {
		m[strconv.FormatFloat(float64(s.Buc), 'f', 5, 32)] = fmt.Sprintf("%d", i)
	}

	return m
}

func (i Index) SumX(b float32) int {
	var sum int

	for _, s := range i.Sta {
		if s.Buc != b {
			sum += s.Cou
		}
	}

	return sum
}
