package win

func sgn(b float32) string {
	if b == 0 {
		return "zer"
	} else if b < 0 {
		return "neg"
	} else {
		return "pos"
	}
}
