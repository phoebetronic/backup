package win

import (
	"encoding/json"
	"path/filepath"

	"github.com/phoebetron/backup/pkg/ind"
)

func (r *run) ind(p string, l ind.List) {
	var err error

	var i ind.Index
	{
		i = ind.Read(p)
	}

	{
		i.Lis = append(i.Lis, l...)
		i.Lis = i.Lis.Sort()
		i.Sta = ind.StaFroMap(i.Lis.Buck())
	}

	var b []byte
	{
		b, err = json.MarshalIndent(i, "", "  ")
		if err != nil {
			panic(err)
		}
	}

	{
		r.fil(filepath.Join(p, "ind", "ind.json"), b)
	}
}
