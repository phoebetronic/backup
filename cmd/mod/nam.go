package mod

import (
	"fmt"
	"math"
)

func nam(b float32, s ...string) string {
	var a string
	if len(s) == 1 {
		a = s[0] + "."
	}

	return fmt.Sprintf("%s.%.3f.%scsv", sgn(b), math.Abs(float64(b)), a)
}
