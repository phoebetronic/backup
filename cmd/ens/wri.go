package ens

import (
	"bytes"
	"encoding/csv"
)

func (r *run) wri(str [][]string, pat string) {
	byt := bytes.NewBuffer([]byte{})
	wri := csv.NewWriter(byt)

	for _, s := range str {
		err := wri.Write(s)
		if err != nil {
			panic(err)
		}
	}

	{
		wri.Flush()
		err := wri.Error()
		if err != nil {
			panic(err)
		}
	}

	{
		r.fil(pat, byt.Bytes())
	}
}
