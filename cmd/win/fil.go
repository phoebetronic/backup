package win

import "io/ioutil"

func (r *run) fil(p string, b []byte) {
	err := ioutil.WriteFile(p, b, 0600)
	if err != nil {
		panic(err)
	}
}
