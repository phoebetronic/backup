package win

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/phoebetron/backup/pkg/mis/win"
	"github.com/phoebetron/series/buff"
	"github.com/phoebetron/series/conf"
	"github.com/phoebetron/series/spec"
)

func (r *run) csv(w []win.Window) {
	{
		p := "./dat/csv"

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	var con buff.Conf
	{
		con = buff.Rand(conf.Conf{
			BU: 10,
			SP: 10,
			TH: 10,
		})
	}

	{
		byt, err := json.MarshalIndent(con.Pipe, "  ", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("produced buffer config\n")
		fmt.Printf("\n")
		fmt.Printf("  %s\n", byt)
		fmt.Printf("\n")
	}

	var flo [][]float32
	for _, v := range w {
		var buf spec.Buff
		{
			buf = buff.With(con)
		}

		for _, t := range v.LE.TR {
			buf.Buff(t)
		}

		{
			flo = append(flo, append(buf.Diff(), v.SI))
		}
	}

	var le1 int
	{
		le1 = len(flo[0])
	}

	var str [][]string
	for _, v := range flo {
		var le2 int
		{
			le2 = len(v)
		}

		if le1 != le2 {
			continue
		}

		var s []string

		for _, f := range v {
			s = append(s, fmt.Sprintf("%.5f", f))
		}

		str = append(str, s)
	}

	byt := bytes.NewBuffer([]byte{})

	see := map[string]struct{}{}
	wri := csv.NewWriter(byt)
	for _, s := range str {
		var k string
		{
			k = strings.Join(s, ",")
		}

		{
			_, exi := see[k]
			if exi {
				continue
			}
			see[k] = struct{}{}
		}

		{
			err := wri.Write(s)
			if err != nil {
				panic(err)
			}
		}
	}

	{
		fmt.Printf("buffered %d window frames in .csv file\n", len(see))
	}

	{
		wri.Flush()
		err := wri.Error()
		if err != nil {
			panic(err)
		}
	}

	{
		pat := filepath.Join("dat", "csv", "win.csv")
		err := ioutil.WriteFile(pat, byt.Bytes(), 0600)
		if err != nil {
			panic(err)
		}
	}
}
