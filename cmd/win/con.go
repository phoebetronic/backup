package win

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/phoebetron/series/buff"
	"github.com/phoebetron/series/conf"
)

func (r *run) connew() buff.Conf {
	var con buff.Conf
	{
		con = buff.Rand(conf.Conf{
			BU: 10,
			SP: 10,
			TH: 10,
			MI: 3,
			MA: 6,
		})
	}

	var pat string
	{
		pat = filepath.Join("./dat", con.Hash(), "con")
	}

	{
		r.dir(pat)
	}

	{
		byt, err := json.MarshalIndent(con, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("produced buffer config\n")
		fmt.Printf("\n")
		fmt.Printf("%s\n", byt)
		fmt.Printf("\n")

		r.fil(filepath.Join(pat, "con.json"), byt)
	}

	return con
}
