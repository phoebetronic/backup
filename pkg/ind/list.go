package ind

import (
	"math/rand"
	"sort"
	"time"
)

type List []Item

func (l List) Buck() map[float32]int {
	m := map[float32]int{}

	for _, i := range l {
		m[i.Buc] += i.Cou
	}

	return m
}

func (l List) Shfl() List {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(l), func(i, j int) { l[i], l[j] = l[j], l[i] })

	return l
}

func (l List) Sort() List {
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Buc > l[j].Buc
	})

	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Tim.Unix() < l[j].Tim.Unix()
	})

	return l
}
