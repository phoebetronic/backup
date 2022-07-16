package win

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/phoebetron/backup/pkg/mis/win"
)

func (r *run) dra(w []win.Window) {
	{
		p := "./dat/dra"

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	var max int
	{
		max = 10
	}

	for i := r.cmdfla.Ind; i < r.cmdfla.Ind+max; i++ {
		var byt []byte
		{
			byt = ren(i, w[i].IN, w[i].LE.PR().FL, w[i].RI.PR().FL)
		}

		{
			pat := filepath.Join("dat", "dra", strconv.Itoa(i)+".svg")
			err := ioutil.WriteFile(pat, byt, 0600)
			if err != nil {
				panic(err)
			}
		}
	}
}
