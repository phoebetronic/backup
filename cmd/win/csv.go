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

const (
	fratra = 0.70
	frates = 0.15
	fraval = 0.15
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
			flo = append(flo, append(buf.Diff(), v.BD, v.TD))
		}
	}

	var le1 int
	{
		le1 = len(flo[0])
	}

	see := map[string]struct{}{}

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

		str = append(str, s)
	}

	{
		a := 0
		b := int(float64(len(str)) * (fratra))
		c := int(float64(len(str)) * (fratra + frates))
		d := int(float64(len(str)) * (fratra + frates + fraval))

		r.wri(str[a:b], "tra")
		r.wri(str[b:c], "tes")
		r.wri(str[c:d], "val")
	}
}

func (r *run) wri(str [][]string, des string) {
	{
		fmt.Printf("buffered %d window frames into %s.csv\n", len(str), des)
	}

	byt := bytes.NewBuffer([]byte{})
	wri := csv.NewWriter(byt)

	{
		var s []string

		for i := 0; i < len(str[0])-2; i++ {
			s = append(s, fmt.Sprintf("%02d", i))
		}

		s = append(s, "BD", "TD")

		err := wri.Write(s)
		if err != nil {
			panic(err)
		}
	}

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
		pat := filepath.Join("dat", "csv", des+".csv")
		err := ioutil.WriteFile(pat, byt.Bytes(), 0600)
		if err != nil {
			panic(err)
		}
	}
}
