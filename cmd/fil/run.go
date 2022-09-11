package fil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/trades/sto/tradesredis"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
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
		r.client = apicliaws.New()
		r.storage = r.newsto()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	{
		sta = time.Date(
			r.flags.Sta.Year(),
			r.flags.Sta.Month(),
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		)
	}

	var end time.Time
	{
		end = sta.AddDate(0, 1, 0)
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("checking backup between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	var tra *trades.Trades
	{
		tra, err = r.storage.Search(sta)
		if tradesredis.IsNotFound(err) {
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

			tra = &trades.Trades{}
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
		} else if err != nil {
			panic(err)
		}
	}

	{
		fmt.Printf("creating frames with %s resolution\n", r.flags.Dur)
	}

	var fra *trades.Framer
	{
		fra = tra.Frame(framer.Config{
			Sta: r.flags.Sta,
			End: r.flags.End,
			Dur: r.flags.Dur,
		})
	}

	var lis []*trades.Trades
	for !fra.Last() {
		lis = append(lis, fra.Next())
	}

	var byt []byte
	{
		byt, err = json.Marshal(lis)
		if err != nil {
			panic(err)
		}
	}

	{
		err := ioutil.WriteFile(r.flags.Pat, byt, 0644)
		if err != nil {
			panic(err)
		}
	}
}
