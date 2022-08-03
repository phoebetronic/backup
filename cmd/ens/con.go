package ens

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/phoebetron/backup/pkg/fil"
	"github.com/phoebetron/series/buff"
)

func (r *run) conall() map[string]bufctx {
	var err error

	all := map[string]bufctx{}

	err = filepath.Walk("dat", func(pat string, inf fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if inf.Name() == "dat" {
			return nil
		}

		if !fil.Exists(filepath.Join(pat, "mod", "zer.0.000.ubj")) {
			return filepath.SkipDir
		}

		var byt []byte
		{
			byt, err = ioutil.ReadFile(filepath.Join(pat, "con", "con.json"))
			if err != nil {
				panic(err)
			}
		}

		var con buff.Conf
		{
			err = json.Unmarshal(byt, &con)
			if err != nil {
				panic(err)
			}
		}

		{
			all[inf.Name()] = bufctx{
				con: con,
			}
		}

		if inf.IsDir() {
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return all
}
