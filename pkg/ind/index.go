package ind

type Index struct {
	Lis List   `json:"lis"`
	Sta []Stat `json:"sta"`
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
