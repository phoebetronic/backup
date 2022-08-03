package ens

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/series/buck"
	"github.com/phoebetron/series/buff"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
)

const (
	fratra = 0.70
	frates = 0.15
	fraval = 0.15
)

type bufctx struct {
	con buff.Conf
	len int
}

type run struct {
	cliaws *apicliaws.AWS
	cmdfla *fla
	misctx map[string]bufctx
	misfra framer.Frames
	stotra trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.cmdfla.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.cliaws = apicliaws.Default()
	}

	{
		r.misctx = r.conall()
	}

	{
		r.misfra = r.franew()
	}

	{
		r.stotra = tradesredis.Default()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	var end time.Time
	{
		sta = r.misfra.Min().Sta
		end = r.misfra.Max().End
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("fetching trades between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	var tra *trades.Trades
	{
		tra, err = r.stotra.Search(sta)
		if err != nil {
			panic(err)
		}
	}

	var fra []*trades.Trades
	{
		fra = tra.Frame(r.misfra)
	}

	var con []buff.Conf
	for h, c := range r.misctx {
		{
			c.len = r.num(fra[0:24], c.con)
			r.misctx[h] = c
		}

		{
			con = append(con, c.con)
		}

		{
			r.dir(filepath.Join("dat", h, "ens"))
		}
	}

	{
		fmt.Printf("creating full CSV files for %d buffer configs\n", len(con))
	}

L0:
	for _, f := range fra {
		var m map[string][]buck.Wndw
		{
			m = buck.Buck(f.TR, buck.Prec(), time.Hour, con...)
		}

		t := map[string][][]string{}

		for h, w := range m {
			for _, v := range w {
				if len(v.SE) == r.misctx[h].len {
					t[h] = append(t[h], r.enc(v.SE))
				} else {
					fmt.Printf(
						"ignoring window of length %d due to strict requirement of window length %d\n",
						len(v.SE),
						r.misctx[h].len,
					)

					{
						continue L0
					}
				}
			}
		}

		for h, l := range t {
			var p string
			{
				p = filepath.Join("dat", h, "ens", "ful.csv")
			}

			{
				r.app(p, l)
			}
		}

		{
			fmt.Printf(
				"added window frames between %s and %s to full CSV files\n",
				scrfmt(f.ST.AsTime()),
				scrfmt(f.EN.AsTime()),
			)
		}
	}

	for h := range r.misctx {
		{
			fmt.Printf("splitting CSV files for buffer config %s\n", h)
		}

		var p string
		{
			p = filepath.Join("dat", h, "ens", "ful.csv")
		}

		var l [][]string
		{
			f, err := ioutil.ReadFile(p)
			if err != nil {
				panic(err)
			}

			l, err = csv.NewReader(bytes.NewReader(f)).ReadAll()
			if err != nil {
				panic(fmt.Sprintf("%s - %s", p, err.Error()))
			}
		}

		{
			w := 0
			x := int(float64(len(l)) * (fratra))
			y := int(float64(len(l)) * (fratra + frates))
			z := int(float64(len(l)) * (fratra + frates + fraval))

			r.wri(l[w:x], filepath.Join("dat", h, "ens", "tra.csv"))
			r.wri(l[x:y], filepath.Join("dat", h, "ens", "tes.csv"))
			r.wri(l[y:z], filepath.Join("dat", h, "ens", "val.csv"))
		}

		{
			err := os.Remove(p)
			if err != nil {
				panic(err)
			}
		}
	}
}
