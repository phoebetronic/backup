package ind

import "sort"

type Stat struct {
	Buc float32 `json:"buc"`
	Cou int     `json:"cou"`
}

func StaFroMap(m map[float32]int) []Stat {
	var s []Stat

	for b, c := range m {
		s = append(s, Stat{Buc: b, Cou: c})
	}

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].Buc > s[j].Buc
	})

	return s
}
