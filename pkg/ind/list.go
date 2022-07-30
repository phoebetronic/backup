package ind

import "sort"

type List []Item

func (l List) Sort() List {
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Buc < l[j].Buc
	})

	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Tim.Unix() < l[j].Tim.Unix()
	})

	return l
}
