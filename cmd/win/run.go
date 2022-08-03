package win

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/series/buff"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
)

type run struct {
	cliaws *apicliaws.AWS
	cmdfla *fla
	miscon buff.Conf
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
		r.miscon = r.connew()
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

	var pat string
	{
		pat = filepath.Join("dat", r.miscon.Hash())
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

	// Since the nature of our buffer implementations is rather dynamic, we need
	// to know how many financial features we are supposed to produce. We take
	// all trades of the first 24 hours and pump them through the buffer
	// implementation created using our desired buffer config in order to get
	// the most probable requirement amount of features produced within each
	// window frame of 1 hour.
	var num int
	{
		num = r.num(fra[0:24])
	}

	for _, f := range fra {
		r.raw(pat, num, f)
	}

	{
		r.ful(pat)
	}

	{
		err := os.RemoveAll(filepath.Join(pat, "raw"))
		if err != nil {
			panic(err)
		}
	}

	{
		r.dir(filepath.Join(pat, "mod"))
	}
}
