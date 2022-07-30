package win

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/phoebetron/backup/pkg/ind"
)

func (r *run) ind(p string, l ind.List) {
	var err error

	var cur []byte
	{
		cur, err = ioutil.ReadFile(p)
		if os.IsNotExist(err) {
			cur = []byte(string("[]"))
		} else if err != nil {
			panic(err)
		}
	}

	var lis ind.List
	{
		err = json.Unmarshal(cur, &lis)
		if err != nil {
			panic(err)
		}
	}

	{
		lis = append(lis, l...)
	}

	var byt []byte
	{
		byt, err = json.MarshalIndent(lis.Sort(), "", "  ")
		if err != nil {
			panic(err)
		}
	}

	{
		r.fil(p, byt)
	}
}
