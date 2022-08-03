package ind

import (
	"fmt"
	"strconv"
)

var (
	buc = []float32{
		+0.007,
		+0.006,
		+0.005,
		+0.004,
		+0.003,
		0,
		-0.003,
		-0.004,
		-0.005,
		-0.006,
		-0.007,
	}
)

func EncI() map[string]string {
	m := map[string]string{}

	for i, b := range buc {
		m[strconv.FormatFloat(float64(b), 'f', 5, 32)] = fmt.Sprintf("%d", i)
	}

	return m
}
