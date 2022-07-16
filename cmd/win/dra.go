package win

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/phoebetron/backup/pkg/mis/win"
	"github.com/phoebetron/trades/typ/trades"
)

func (r *run) dra(w []win.Window) {
	{
		p := "./dat/"

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
		le := &trades.Trades{
			TR: w[i].LE,
		}
		ri := &trades.Trades{
			TR: w[i].RI,
		}

		var byt []byte
		{
			byt = ren(i, w[i].IN, le.PR().FL, ri.PR().FL)
		}

		{
			pat := filepath.Join("dat", strconv.Itoa(i)+".gold.svg")
			err := ioutil.WriteFile(pat, byt, 0600)
			if err != nil {
				panic(err)
			}
		}
	}
}
