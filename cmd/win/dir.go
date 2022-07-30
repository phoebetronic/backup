package win

import "os"

func (r *run) dir(p string) {
	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
