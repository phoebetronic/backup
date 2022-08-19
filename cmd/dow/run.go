package dow

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

type run struct {
	client  *apicliaws.AWS
	flags   *flags
	storage trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.client = apicliaws.Default()
		r.storage = r.newsto()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	{
		sta = r.flags.Time
	}

	var end time.Time
	{
		end = sta.AddDate(0, 1, 0)
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("checking backup between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	var buc string
	{
		buc = "xh3b4sd-phoebe-backup"
	}

	var pre string
	{
		pre = fmt.Sprintf("tra-raw.exc-%s.ass-%s", r.storage.Market().Exc(), r.storage.Market().Ass())
	}

	var suf string
	{
		suf = fmt.Sprintf("%s.pb.raw", bacfmt(sta))
	}

	var byt []byte
	{
		byt, err = r.client.Download(buc, filepath.Join(pre, suf))
		if err != nil {
			panic(err)
		}
	}

	tra := &trades.Trades{}
	{
		err := proto.Unmarshal(byt, tra)
		if err != nil {
			panic(err)
		}
	}

	{
		err := r.storage.Create(tra.ST.AsTime(), tra)
		if tradesredis.IsAlreadyExists(err) {
			err = r.storage.Update(tra.ST.AsTime(), tra)
			if err != nil {
				panic(err)
			}
		} else if err != nil {
			panic(err)
		}
	}
}
